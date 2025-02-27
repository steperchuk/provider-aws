/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DataSourceConfigurationObservation struct {
}

type DataSourceConfigurationParameters struct {

	// Specifies how the results of an operation will be stored by the caller. Valid values: SingleUse, Storage. Default: SingleUse.
	// +kubebuilder:validation:Optional
	IntendedUse *string `json:"intendedUse,omitempty" tf:"intended_use,omitempty"`
}

type PlaceIndexObservation struct {

	// The timestamp for when the place index resource was created in ISO 8601 format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) for the place index resource. Used to specify a resource across AWS.
	IndexArn *string `json:"indexArn,omitempty" tf:"index_arn,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The timestamp for when the place index resource was last update in ISO 8601.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type PlaceIndexParameters struct {

	// Specifies the geospatial data provider for the new place index.
	// +kubebuilder:validation:Required
	DataSource *string `json:"dataSource" tf:"data_source,omitempty"`

	// Configuration block with the data storage option chosen for requesting Places. Detailed below.
	// +kubebuilder:validation:Optional
	DataSourceConfiguration []DataSourceConfigurationParameters `json:"dataSourceConfiguration,omitempty" tf:"data_source_configuration,omitempty"`

	// The optional description for the place index resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PlaceIndexSpec defines the desired state of PlaceIndex
type PlaceIndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlaceIndexParameters `json:"forProvider"`
}

// PlaceIndexStatus defines the observed state of PlaceIndex.
type PlaceIndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlaceIndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PlaceIndex is the Schema for the PlaceIndexs API. Provides a Location Service Place Index.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PlaceIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlaceIndexSpec   `json:"spec"`
	Status            PlaceIndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlaceIndexList contains a list of PlaceIndexs
type PlaceIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlaceIndex `json:"items"`
}

// Repository type metadata.
var (
	PlaceIndex_Kind             = "PlaceIndex"
	PlaceIndex_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PlaceIndex_Kind}.String()
	PlaceIndex_KindAPIVersion   = PlaceIndex_Kind + "." + CRDGroupVersion.String()
	PlaceIndex_GroupVersionKind = CRDGroupVersion.WithKind(PlaceIndex_Kind)
)

func init() {
	SchemeBuilder.Register(&PlaceIndex{}, &PlaceIndexList{})
}
