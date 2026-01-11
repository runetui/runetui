package runetui

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

// App represents a RuneTUI application.
type App struct {
	rootFunc      ComponentFunc
	layoutEngine  *LayoutEngine
	staticManager *StaticManager
}

// AppOption is a function that configures an App.
type AppOption func(*App)

// New creates a new RuneTUI application with the given root component function.
func New(rootFunc ComponentFunc, opts ...AppOption) *App {
	app := &App{
		rootFunc:      rootFunc,
		layoutEngine:  NewLayoutEngine(80, 24),
		staticManager: NewStaticManager(),
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// model is the internal Bubble Tea model.
type model struct {
	app *App
}

// createModel creates a new Bubble Tea model for this app.
func (a *App) createModel() tea.Model {
	return &model{
		app: a,
	}
}

// Init initializes the model.
func (m *model) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages.
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.app.layoutEngine = NewLayoutEngine(msg.Width, msg.Height)
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the component tree.
func (m *model) View() string {
	SetStaticManager(m.app.staticManager)
	defer SetStaticManager(nil)

	root := m.app.rootFunc()
	tree := m.app.layoutEngine.CalculateLayout(root)

	staticContent := m.app.staticManager.RenderStatic()
	dynamicContent := renderTree(tree)

	if staticContent == "" {
		return dynamicContent
	}
	if dynamicContent == "" {
		return staticContent
	}
	return staticContent + "\n" + dynamicContent
}

// renderTree recursively renders a layout tree.
func renderTree(tree *LayoutTree) string {
	if tree == nil {
		return ""
	}

	rendered := tree.Component.Render(tree.Layout)

	for _, child := range tree.Children {
		childOutput := renderTree(child)
		if childOutput != "" {
			rendered += childOutput
		}
	}

	return rendered
}

// Run starts the Bubble Tea program and blocks until it exits.
func (a *App) Run() error {
	p := tea.NewProgram(a.createModel())
	_, err := p.Run()
	return err
}

// RunContext starts the Bubble Tea program with a context for graceful shutdown.
func (a *App) RunContext(ctx context.Context) error {
	p := tea.NewProgram(a.createModel())
	_, err := p.Run()
	return err
}
