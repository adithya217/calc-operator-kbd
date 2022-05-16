package services

import (
	v1 "calc-operator/api/v1"
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type resourceService struct {
	logger  logr.Logger
	client  client.Client
	context context.Context
}

type IResourceService interface {
	UpdateStatus(v1.Calculator) error
}

func NewResourceService(logger logr.Logger, client client.Client, ctx context.Context) IResourceService {
	return &resourceService{
		logger:  logger,
		client:  client,
		context: ctx,
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
