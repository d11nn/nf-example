package processor

import (
	"fmt"
	"net/http"
	"sort"

	nf_context "github.com/Alonza0314/nf-example/internal/context"
	"github.com/gin-gonic/gin"
)

func (p *Processor) ListUes(c *gin.Context) {
	ctx := p.Context()
	ctx.UeMu.RLock()
	defer ctx.UeMu.RUnlock()

	ues := make([]nf_context.UeInfo, 0, len(ctx.UeData))
	for _, ue := range ctx.UeData {
		ues = append(ues, ue)
	}
	sort.Slice(ues, func(i, j int) bool {
		return ues[i].Supi < ues[j].Supi
	})
	c.JSON(http.StatusOK, ues)
}

func (p *Processor) GetUe(c *gin.Context, supi string) {
	ctx := p.Context()
	ctx.UeMu.RLock()
	defer ctx.UeMu.RUnlock()

	if ue, ok := ctx.UeData[supi]; ok {
		c.JSON(http.StatusOK, ue)
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not registered", supi))
}

func (p *Processor) RegisterUe(c *gin.Context, ue nf_context.UeInfo) {
	ctx := p.Context()
	ctx.UeMu.Lock()
	defer ctx.UeMu.Unlock()

	if _, ok := ctx.UeData[ue.Supi]; ok {
		c.String(http.StatusConflict, fmt.Sprintf("[%s] already registered", ue.Supi))
		return
	}
	ctx.UeData[ue.Supi] = ue
	c.JSON(http.StatusCreated, ue)
}
