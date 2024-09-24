package flow

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	// GroupName is the name of the API group this package's resources belong to.
	GroupName = "flow.typhoon.zeiss.com"
)

var (
	// JQTransformationResource respresents a JQ transformation.
	JQTransformationResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "jqtransformations",
	}

	// WorkerTransformationResource respresents a Worker transformation.
	WorkerTransformationResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "workertransformations",
	}

	// SynchronizerResource respresents a Synchronizer.
	SynchronizerResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "synchronizers",
	}

	// TransformationResource respresents a Bumblebee transformation.
	TransformationResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "transformations",
	}

	// XMLToJSONTransformationResource respresents a XML to JSON transformation.
	XMLToJSONTransformationResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "xmltojsontansformations",
	}

	// XSLTTransformationResource respresents a XSLT transformation.
	XSLTTransformationResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "xslttransformations",
	}

	// BridgeResource respresents a bundled application flow.
	BridgeResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "bridges",
	}
)
