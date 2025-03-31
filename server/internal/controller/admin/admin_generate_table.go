package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/api/admin/generate"
	"server/internal/service"
)

func (c *ControllerGenerate) GetTables(ctx context.Context, req *generate.GetTablesReq) (res *generate.GetTablesRes, err error) {
	res, err = service.Generate().GetTables(ctx, req)
	return
}
func (c *ControllerGenerate) GetTableColumns(ctx context.Context, req *generate.GetTableColumnsReq) (res *generate.GetTableColumnsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerGenerate) GenerateSql(ctx context.Context, req *generate.GenerateSqlReq) (res *generate.GenerateSqlRes, err error) {
	res, err = service.Generate().GenerateSql(ctx, req)
	return
}
func (c *ControllerGenerate) ExecuteSql(ctx context.Context, req *generate.ExecuteSqlReq) (res *generate.ExecuteSqlRes, err error) {
	err = service.Generate().ExecuteSql(ctx, req)
	return
}
