package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_project/internal/logic"
	"go_project/internal/svc"
	"go_project/internal/types"
)

func GetCourseListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCourseListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCourseListLogic(r.Context(), svcCtx)
		resp, err := l.GetCourseList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
