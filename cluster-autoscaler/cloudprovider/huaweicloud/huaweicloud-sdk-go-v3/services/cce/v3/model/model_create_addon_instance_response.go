package model

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAddonInstanceResponse struct {
	// API类型，固定值“Addon”，该值不可修改。

	Kind *string `json:"kind,omitempty"`
	// API版本，固定值“v3”，该值不可修改。

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *InstanceSpec `json:"spec,omitempty"`

	Status         *AddonInstanceStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateAddonInstanceResponse", string(data)}, " ")
}
