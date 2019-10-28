package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAwsEMRPublicAccessBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsEmrPublicAccessBlockCreate,
		Read:   resourceAwsEmrPublicAccessBlockRead,
		Delete: resourceAwsEmrPublicAccessBlockDelete,

		Schema: map[string]*schema.Schema{
			"block_public_security_group_rules": {
				Type:          schema.TypeBool,
				Required:      true,
				Default        true,
			},
			"permitted_public_security_group_rule_range": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_range": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_range": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAwsEmrPublicAccessBlockCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).emrconn
	var portRanges := []*PortRange{}
	// TBD ここをいい感じに書く
	for rule := ; range d.Get("permitted_public_security_group_rule_range") {
		portRanges = append(portRanges, &emr.PortRange{
			MaxRange: 1,
			MinRange: 2,
		})
	}
	
	resp, err := conn.PutBlockPublicAccessConfiguration(
		&emr.PutBlockPublicAccessConfigurationInput{
			BlockPublicAccessConfiguration: &emr.BlockPublicAccessConfiguration{
				BlockPublicSecurityGroupRules: aws.Bool(d.Get("block_public_security_group_rules").(bool)),
				PermittedPublicSecurityGroupRuleRanges: &portRanges
			} 
		}
	)

	if err != nil {
		return err
	}
	return resourceAwsEmrPublicAccessBlockRead(d, meta)
}

func resourceAwsEmrPublicAccessBlockRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsEmrPublicAccessBlockDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).emrconn

	resp, err := conn.PutBlockPublicAccessConfiguration(
		&emr.PutBlockPublicAccessConfigurationInput{
			BlockPublicAccessConfiguration: &emr.BlockPublicAccessConfiguration{
				BlockPublicSecurityGroupRules: aws.Bool(false)),
				PermittedPublicSecurityGroupRuleRanges: []*PortRange{
					&emr.PortRange{
						// Rollback to Default 
						MaxRange: 22,
						MinRange: 22,
					}
				}
			} 
		}
	)

	if err != nil {
		return err
	}
	return nil
}
