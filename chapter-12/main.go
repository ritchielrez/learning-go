package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    choices []string // items on the todo list
    cursor int       // which todo list item our cursor is pointing at

    // This is actually a set
    selected map[int]struct{} // which todo iterms are selected
}

func initialModel() model {
    return model {
        choices: []string{"Buy appls", "Buy oranges", "Buy anything else"},

        // A map which indicates which choices are selected. We're using
        // the map like a mathematical set. The keys refer to the indexes
        // of the `choices` slice, above.
        selected: make(map[int]struct{}),
    }
}

// Think the `model` as a class and `Init()` as a method of it
func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    // Return updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := "What should we buy from the market?\n\n"

    // Iterate over our choices
    for i, choice := range m.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "X" // selected
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice);
    }

    // The footer
    s += "\nPress q to quit\n"

    // Send the UI for rendering
    return s
}

func main() {
    p := tea.NewProgram(initialModel())

    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v\n", err)
        os.Exit(1)
    }
}
