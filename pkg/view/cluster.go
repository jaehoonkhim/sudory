package view

import (
	"net/http"

	"github.com/NexClipper/sudory-prototype-r1/pkg/control/operator"
	"github.com/NexClipper/sudory-prototype-r1/pkg/model"
	"github.com/labstack/echo/v4"
)

type CreateCluster struct {
	opr *operator.Cluster
}

func NewCreateCluster(o operator.Operator) Viewer {
	return &CreateCluster{opr: o.(*operator.Cluster)}
}

func (v *CreateCluster) fromModel(m *model.ReqCluster) {
	v.opr.Name = m.Name
	v.opr.Response = v.Response
}

func (v *CreateCluster) Request(ctx echo.Context) error {
	reqModel := &model.ReqCluster{}
	if err := ctx.Bind(reqModel); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	v.fromModel(reqModel)
	if err := v.opr.Create(ctx); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return nil
}

func (v *CreateCluster) Response(ctx echo.Context, m model.Modeler) error {
	if err := ctx.JSON(http.StatusOK, nil); err != nil {
		return err
	}
	return nil
}

type GetCluster struct {
	opr *operator.Cluster
}

func NewGetCluster(o operator.Operator) Viewer {
	return &GetCluster{opr: o.(*operator.Cluster)}
}

func (v *GetCluster) fromModel(id string) {
	v.opr.ID = id
	v.opr.Response = v.Response
}

func (v *GetCluster) Request(ctx echo.Context) error {
	id := ctx.Param("id")

	v.fromModel(id)
	if err := v.opr.Get(ctx); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return nil
}

func (v *GetCluster) Response(ctx echo.Context, m model.Modeler) error {
	cluster := m.(*model.Cluster)
	return ctx.JSON(http.StatusOK, cluster)
}
