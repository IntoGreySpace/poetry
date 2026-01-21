// Copyright (c) 2025, The Grey-Space Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/icons"
	_ "cogentcore.org/core/text/tex" // include this to get math
	"cogentcore.org/core/tree"
	_ "cogentcore.org/lab/yaegilab"
)

//go:embed content
var econtent embed.FS

//go:embed icon.svg
var icon string

func main() {
	core.AppIcon = icon
	content.Settings.SiteTitle = "The Grey Space"
	content.OfflineURL = "https://poetry.grey-space.org"
	b := core.NewBody(content.Settings.SiteTitle)
	ct := content.NewContent(b).SetContent(econtent)
	ctx := ct.Context
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(func(p *tree.Plan) {
			ct.MakeToolbar(p)
			ct.MakeToolbarPDF(p)
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://github.com/IntoGreySpace/poetry")
				w.SetText("GitHub").SetIcon(icons.GitHub)
			})
		})
	})

	b.RunMainWindow()
}
