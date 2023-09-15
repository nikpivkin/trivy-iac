package cloudwatch

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/cloudwatch"
	"github.com/nikpivkin/trivy-iac/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) cloudwatch.CloudWatch {
	return cloudwatch.CloudWatch{
		LogGroups: getLogGroups(cfFile),
		Alarms:    nil,
	}
}
