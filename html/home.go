package html

import (
	"time"

	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"

	"app/model"
)

// HomePage is the front page of the app.
func HomePage(props PageProps, things []model.Thing, now time.Time) g.Node {
	props.Title = "Home"

	return page(props,
		h.Div(h.Class("prose prose-indigo prose-lg md:prose-xl"),
			h.H1(g.Text("Welcome to the gomponents starter kit")),

			h.P(g.Text("It uses gomponents, HTMX, and Tailwind CSS, and you can use it as a template for your new app. 😎")),

			h.P(h.A(h.Href("https://github.com/maragudk/gomponents-starter-kit"), g.Text("See gomponents-starter-kit on GitHub"))),

			h.H2(g.Text("Try HTMX")),

			h.Button(
				h.Class("rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"),
				g.Text("Get things with HTMX"), hx.Get("/"), hx.Target("#things")),

			h.Div(h.ID("things"),
				ThingsPartial(things, now),
			),
		),
	)
}

// ThingsPartial is a partial for showing a list of things, returned directly if the request is an HTMX request,
// and used in [HomePage].
func ThingsPartial(things []model.Thing, now time.Time) g.Node {
	return g.Group{
		h.P(g.Textf("Here are %v things from the mock database (updated %v):", len(things), now.Format(time.TimeOnly))),
		h.Ul(
			g.Map(things, func(t model.Thing) g.Node {
				return h.Li(g.Text(t.Name))
			}),
		),
	}
}
