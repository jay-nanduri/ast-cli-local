package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	resultsReader "github.com/checkmarxDev/sast-results/pkg/reader"
	resultsHelpers "github.com/checkmarxDev/sast-results/pkg/web/helpers"
	resultsRaw "github.com/checkmarxDev/sast-results/pkg/web/path/raw"

	commonParams "github.com/checkmarxDev/ast-cli/internal/params"

	"github.com/checkmarxDev/ast-cli/internal/wrappers"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	failedListingResults = "Failed listing results"
)

var (
	filterResultsListFlagUsage = fmt.Sprintf("Filter the list of results. Use ';' as the delimeter for arrays. Available filters are: %s",
		strings.Join([]string{
			commonParams.ScanIDQueryParam,
			commonParams.LimitQueryParam,
			commonParams.OffsetQueryParam,
			commonParams.SortQueryParam,
			commonParams.IncludeNodesQueryParam,
			commonParams.NodeIDsQueryParam,
			commonParams.QueryQueryParam,
			commonParams.GroupQueryParam,
			commonParams.StatusQueryParam,
			commonParams.SeverityQueryParam}, ","))
)

func NewResultCommand(resultsWrapper wrappers.ResultsWrapper) *cobra.Command {
	resultCmd := &cobra.Command{
		Use:   "result",
		Short: "Retrieve results",
	}

	listResultsCmd := &cobra.Command{
		Use:   "list <scan-id>",
		Short: "List results for a given scan",
		RunE:  runGetResultByScanIDCommand(resultsWrapper),
	}
	listResultsCmd.PersistentFlags().StringSlice(filterFlag, []string{}, filterResultsListFlagUsage)
	addFormatFlag(listResultsCmd, formatList, formatJSON)
	resultCmd.AddCommand(listResultsCmd)
	return resultCmd
}

func runGetResultByScanIDCommand(resultsWrapper wrappers.ResultsWrapper) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var resultResponseModel *resultsRaw.ResultsCollection
		var errorModel *resultsHelpers.WebError
		var err error
		if len(args) == 0 {
			return errors.Errorf("%s: Please provide a scan ID", failedListingResults)
		}

		scanID := args[0]
		params, err := getFilters(cmd)
		if err != nil {
			return errors.Wrapf(err, "%s", failedListingResults)
		}
		params[commonParams.ScanIDQueryParam] = scanID

		resultResponseModel, errorModel, err = resultsWrapper.GetByScanID(params)
		if err != nil {
			return errors.Wrapf(err, "%s", failedListingResults)
		}
		// Checking the response
		if errorModel != nil {
			return errors.Errorf("%s: CODE: %d, %s", failedListingResults, errorModel.Code, errorModel.Message)
		} else if resultResponseModel != nil {
			f, _ := cmd.Flags().GetString(formatFlag)
			if IsFormat(f, formatJSON) {
				var resultsJSON []byte
				resultsJSON, err = json.Marshal(resultResponseModel)
				if err != nil {
					return errors.Wrapf(err, "%s: failed to serialize results response ", failedGettingAll)
				}

				fmt.Fprintln(cmd.OutOrStdout(), string(resultsJSON))
				return nil
			}

			// Not supporting table view because it gets ugly
			return outputResultsPretty(resultResponseModel.Results)
		}

		return nil
	}
}

func outputResultsPretty(results []*resultsReader.Result) error {
	fmt.Println("************ Results ************")
	for i := 0; i < len(results); i++ {
		outputSingleResult(&resultsReader.Result{
			ResultQuery: resultsReader.ResultQuery{
				QueryID:   results[i].QueryID,
				QueryName: results[i].QueryName,
				Severity:  results[i].Severity,
				CweID:     results[i].CweID,
			},
			SimilarityID: results[i].SimilarityID,
			UniqueID:     results[i].UniqueID,
			FirstScanID:  results[i].FirstScanID,
			FirstFoundAt: results[i].FirstFoundAt,
			FoundAt:      results[i].FoundAt,
			Status:       results[i].Status,
			PathSystemID: results[i].PathSystemID,
			Nodes:        results[i].Nodes,
		})
		fmt.Println()
	}
	return nil
}

func outputSingleResult(model *resultsReader.Result) {
	fmt.Println("Result Unique ID:", model.UniqueID)
	fmt.Println("Query ID:", model.QueryID)
	fmt.Println("Query Name:", model.QueryName)
	fmt.Println("Severity:", model.Severity)
	fmt.Println("CWE ID:", model.CweID)
	fmt.Println("Similarity ID:", model.SimilarityID)
	fmt.Println("First Scan ID:", model.FirstScanID)
	fmt.Println("Found At:", model.FoundAt)
	fmt.Println("First Found At:", model.FirstFoundAt)
	fmt.Println("Status:", model.Status)
	fmt.Println("Path System ID:", model.PathSystemID)
	fmt.Println()
	fmt.Println("************ Nodes ************")
	for i := 0; i < len(model.Nodes); i++ {
		outputSingleResultNodePretty(model.Nodes[i])
		fmt.Println()
	}
}
