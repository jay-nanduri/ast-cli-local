package wrappers

import (
	resultsReader "github.com/checkmarxDev/sast-results/pkg/reader"
	resultsHelpers "github.com/checkmarxDev/sast-results/pkg/web/helpers"
	resultsRaw "github.com/checkmarxDev/sast-results/pkg/web/path/raw"
)

type ResultsMockWrapper struct{}

func (r ResultsMockWrapper) GetByScanID(_ map[string]string) (*resultsRaw.ResultsCollection, *resultsHelpers.WebError, error) {
	const mock = "MOCK"
	return &resultsRaw.ResultsCollection{
		Results: []*resultsReader.Result{
			{
				ResultQuery: resultsReader.ResultQuery{
					QueryID:   0,
					QueryName: mock,
					Severity:  mock,
				},
				PathSystemID: mock,
				ID:           mock,
				FirstScanID:  mock,
				FirstFoundAt: mock,
				FoundAt:      mock,
				Status:       mock,
			},
		},
		TotalCount: 1,
	}, nil, nil
}
