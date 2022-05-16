package services

import (
	v1 "calc-operator/api/v1"
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type resourceService struct {
	logger   logr.Logger
	client   client.Client
	context  context.Context
	recorder record.EventRecorder
}

type IResourceService interface {
	UpdateStatus(v1.Calculator) error
	RecordEvent(object runtime.Object, eventtype, reason, message string) error
}

func NewResourceService(logger logr.Logger, client client.Client, ctx context.Context, recorder record.EventRecorder) IResourceService {
	return &resourceService{
		logger:   logger,
		client:   client,
		context:  ctx,
		recorder: recorder,
	}
}

func (rs resourceService) UpdateStatus(data v1.Calculator) error {
	rs.logger.Info("about to update status", "value", fmt.Sprintf("%#v", data.Status))

	err := rs.client.Status().Update(rs.context, &data)
	if err != nil {
		rs.logger.Info("failed to update status", "error", err)
	}

	return err
}

func (rs resourceService) RecordEvent(object runtime.Object, eventtype, reason, message string) error {
	rs.recorder.Event(object, eventtype, reason, message)
	return nil
}
