package envoycharts

import (
	// "bytes"
	// "io"
	// "os"

	// "github.com/go-echarts/go-echarts/v2/charts"
	// "github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/go-echarts/go-echarts/v2/render"
	ectpls "github.com/cloudkucooland/EnvoyCharts/templates"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

func useECTemplates() {
	tpls.BaseTpl = ectpls.BaseTpl
	tpls.ChartTpl = ectpls.ChartTpl
	tpls.HeaderTpl = ectpls.HeaderTpl
}

/*
type ecRenderer struct {
    c interface{}
    before []func()
}

func NewECRenderer(c interface{}, before ...func()) render.Renderer {
    return &ecRenderer{c: c, before: before}
}

func (r *ecRenderer) Render(w io.Writer) error {
    for _, fn := range r.before {
        fn()
    }
    contents := []string{templates.HeaderTpl, templates.BaseTpl, templates.ChartTpl}
    template := render.MustTemplate("chart", contents)

    var buf bytes.Buffer
    if err := template.ExecuteTemplate(&buf, "chart", r.c); err != nil {
        return err
    }

    _, err := w.Write(buf.Bytes())
    return err
} */
