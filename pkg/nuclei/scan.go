package nuclei

import (
	"context"
	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog/disk"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog/loader"
	"github.com/projectdiscovery/nuclei/v2/pkg/core"
	"github.com/projectdiscovery/nuclei/v2/pkg/core/inputs"
	"github.com/projectdiscovery/nuclei/v2/pkg/output"
	"github.com/projectdiscovery/nuclei/v2/pkg/parsers"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/contextargs"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/hosterrorscache"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/interactsh"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolinit"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolstate"
	"github.com/projectdiscovery/nuclei/v2/pkg/reporting"
	"github.com/projectdiscovery/nuclei/v2/pkg/testutils"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
	"github.com/projectdiscovery/ratelimit"
	pb "github.com/pyneda/nuclei-api/pkg/service"
	"log"
	"os"
	"path"
	"time"
)

func eventToScanResult(event *output.ResultEvent) *pb.ScanResult {
	var info *pb.ScanResultInfo
	if event.Info.Classification != nil {
		info = &pb.ScanResultInfo{
			Name:        event.Info.Name,
			Description: event.Info.Description,
			Severity:    event.Info.SeverityHolder.Severity.String(),
			Remediation: event.Info.Remediation,
			Tags:        event.Info.Tags.ToSlice(),
			References:  event.Info.Reference.ToSlice(),
			Classification: &pb.ScanResultClassification{
				Cves:       event.Info.Classification.CVEID.ToSlice(),
				Cwes:       event.Info.Classification.CWEID.ToSlice(),
				Cpe:        event.Info.Classification.CPE,
				CvssVector: event.Info.Classification.CVSSMetrics,
				CvssScore:  event.Info.Classification.CVSSScore,
			},
		}
	}

	var interaction *pb.Interaction
	if event.Interaction != nil {
		interaction = &pb.Interaction{
			Protocol:      event.Interaction.Protocol,
			UniqueId:      event.Interaction.UniqueID,
			FullId:        event.Interaction.FullId,
			Qtype:         event.Interaction.QType,
			RawRequest:    []byte(event.Interaction.RawRequest),
			RawResponse:   []byte(event.Interaction.RawResponse),
			SmtpFrom:      event.Interaction.SMTPFrom,
			RemoteAddress: event.Interaction.RemoteAddress,
			Timestamp:     event.Interaction.Timestamp.String(),
		}
	}

	return &pb.ScanResult{
		TemplateId:       event.TemplateID,
		Template:         event.Template,
		Info:             info,
		MatcherName:      event.MatcherName,
		ExtractorName:    event.ExtractorName,
		Type:             event.Type,
		Host:             event.Host,
		Path:             event.Path,
		Matched:          event.Matched,
		ExtractedResults: event.ExtractedResults,
		Request:          []byte(event.Request),
		Response:         []byte(event.Response),
		Ip:               event.IP,
		Timestamp:        event.Timestamp.String(),
		CurlCommand:      event.CURLCommand,
		MatcherStatus:    event.MatcherStatus,
		Interaction:      interaction,
	}
}

func Scan(scanRequest *pb.ScanRequest, stream pb.NucleiApi_ScanServer) {
	cache := hosterrorscache.New(30, hosterrorscache.DefaultMaxHostsCount, nil)
	defer cache.Close()

	mockProgress := &testutils.MockProgressClient{}
	reportingClient, _ := reporting.New(&reporting.Options{}, "")
	defer reportingClient.Close()

	outputWriter := testutils.NewMockOutputWriter()
	outputWriter.WriteCallback = func(event *output.ResultEvent) {
		// printFields(*event)
		log.Printf("Got Result: %v\n", event.TemplateID)
		result := eventToScanResult(event)
		err := stream.Send(result)
		if err != nil {
			log.Printf("Error sending %v result to client: %v", event.TemplateID, err)
		}
	}

	defaultOpts := getDefaultOptions()

	defaultOpts.AutomaticScan = scanRequest.AutomaticScan
	defaultOpts.Headless = scanRequest.Headless
	defaultOpts.NewTemplates = scanRequest.NewTemplates
	assignIfNotEmpty(&defaultOpts.Tags, &scanRequest.Tags)
	assignIfNotEmpty(&defaultOpts.ExcludeTags, &scanRequest.ExcludeTags)
	assignIfNotEmpty(&defaultOpts.Workflows, &scanRequest.Workflows)
	assignIfNotEmpty(&defaultOpts.WorkflowURLs, &scanRequest.WorkflowUrls)
	assignIfNotEmpty(&defaultOpts.Templates, &scanRequest.Templates)
	assignIfNotEmpty(&defaultOpts.TemplateURLs, &scanRequest.TemplateUrls)
	assignIfNotEmpty(&defaultOpts.ExcludedTemplates, &scanRequest.ExcludedTemplates)
	assignIfNotEmpty(&defaultOpts.ExcludeMatchers, &scanRequest.ExcludeMatchers)
	assignIfNotEmpty(&defaultOpts.Authors, &scanRequest.Authors)
	assignIfNotEmpty(&defaultOpts.IncludeTags, &scanRequest.IncludeTags)
	assignIfNotEmpty(&defaultOpts.IncludeTemplates, &scanRequest.IncludeTemplates)
	assignIfNotEmpty(&defaultOpts.IncludeIds, &scanRequest.IncludeIds)
	assignIfNotEmpty(&defaultOpts.ExcludeIds, &scanRequest.ExcludeIds)

	for _, severityString := range scanRequest.Severities {
		convertedSeverity, err := stringToSeverity(severityString)
		if err != nil {
			log.Printf("Invalid severity: %s", severityString)
			continue
		}
		defaultOpts.Severities.Set(convertedSeverity.String())
	}
	for _, severityString := range scanRequest.ExcludeSeverities {
		convertedSeverity, err := stringToSeverity(severityString)
		if err != nil {
			log.Printf("Invalid severity: %s", severityString)
			continue
		}
		defaultOpts.ExcludeSeverities.Set(convertedSeverity.String())
	}
	for _, protocolString := range scanRequest.Protocols {
		convertedProtocol, err := stringToProtocol(protocolString)
		if err != nil {
			log.Printf("Invalid protocol: %s", protocolString)
			continue
		}
		defaultOpts.Protocols.Set(convertedProtocol.String())
	}
	for _, protocolString := range scanRequest.ExcludeProtocols {
		convertedProtocol, err := stringToProtocol(protocolString)
		if err != nil {
			log.Printf("Invalid protocol: %s", protocolString)
			continue
		}
		defaultOpts.ExcludeProtocols.Set(convertedProtocol.String())
	}

	protocolstate.Init(defaultOpts)
	protocolinit.Init(defaultOpts)

	interactOpts := interactsh.DefaultOptions(outputWriter, reportingClient, mockProgress)
	interactClient, err := interactsh.New(interactOpts)
	if err != nil {
		log.Fatalf("Could not create interact client: %s\n", err)
	}
	defer interactClient.Close()

	home, _ := os.UserHomeDir()
	catalog := disk.NewCatalog(path.Join(home, "nuclei-templates"))
	executerOpts := protocols.ExecutorOptions{
		Output:          outputWriter,
		Options:         defaultOpts,
		Progress:        mockProgress,
		Catalog:         catalog,
		IssuesClient:    reportingClient,
		RateLimiter:     ratelimit.New(context.Background(), 150, time.Second),
		Interactsh:      interactClient,
		HostErrorsCache: cache,
		Colorizer:       aurora.NewAurora(true),
		ResumeCfg:       types.NewResumeCfg(),
	}
	engine := core.New(defaultOpts)
	engine.SetExecuterOptions(executerOpts)

	workflowLoader, err := parsers.NewLoader(&executerOpts)
	if err != nil {
		log.Fatalf("Could not create workflow loader: %s\n", err)
	}
	executerOpts.WorkflowLoader = workflowLoader

	store, err := loader.New(loader.NewConfig(defaultOpts, catalog, executerOpts))
	if err != nil {
		log.Fatalf("Could not create loader client: %s\n", err)
	}
	store.Load()

	inputArgs := []*contextargs.MetaInput{}

	for _, target := range scanRequest.GetTargets() {
		inputArgs = append(inputArgs, &contextargs.MetaInput{Input: target})
	}

	input := &inputs.SimpleInputProvider{Inputs: inputArgs}
	_ = engine.Execute(store.Templates(), input)
	engine.WorkPool().Wait() // Wait for the scan to finish
	log.Println("Scan finished")
}
