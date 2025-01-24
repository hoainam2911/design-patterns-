package main

import (
	"fmt"
)

// Command interface defines a method for executing a command
type Command interface {
	Execute()
	Undo()
}

// Editor represents the receiver in the Command pattern
type Editor struct {
	Text     string
	Clipboard string
}

func (e *Editor) GetSelection() string {
	return e.Text // For simplicity, return all text as selection
}

func (e *Editor) DeleteSelection() {
	e.Text = "" // Clear the text
}

func (e *Editor) ReplaceSelection(text string) {
	e.Text = text
}

// Concrete commands

type CopyCommand struct {
	editor *Editor
}

func (c *CopyCommand) Execute() {
	c.editor.Clipboard = c.editor.GetSelection()
	fmt.Println("Copied to clipboard:", c.editor.Clipboard)
}

func (c *CopyCommand) Undo() {
	// No state change to undo for Copy
}

type CutCommand struct {
	editor *Editor
	backup string
}

func (c *CutCommand) Execute() {
	c.backup = c.editor.Text
	c.editor.Clipboard = c.editor.GetSelection()
	c.editor.DeleteSelection()
	fmt.Println("Cut to clipboard:", c.editor.Clipboard)
}

func (c *CutCommand) Undo() {
	c.editor.Text = c.backup
	fmt.Println("Undo Cut. Restored text:", c.editor.Text)
}

type PasteCommand struct {
	editor *Editor
	backup string
}

func (p *PasteCommand) Execute() {
	p.backup = p.editor.Text
	p.editor.ReplaceSelection(p.editor.Clipboard)
	fmt.Println("Pasted text from clipboard:", p.editor.Text)
}

func (p *PasteCommand) Undo() {
	p.editor.Text = p.backup
	fmt.Println("Undo Paste. Restored text:", p.editor.Text)
}

// CommandHistory maintains a stack of executed commands for undo functionality
type CommandHistory struct {
	history []Command
}

func (h *CommandHistory) Push(cmd Command) {
	h.history = append(h.history, cmd)
}

func (h *CommandHistory) Pop() Command {
	if len(h.history) == 0 {
		return nil
	}
	cmd := h.history[len(h.history)-1]
	h.history = h.history[:len(h.history)-1]
	return cmd
}

// Application acts as the invoker
func main() {
	editor := &Editor{Text: "Hello, World!"}
	history := &CommandHistory{}

	fmt.Println("Initial text:", editor.Text)

	// Execute Copy Command
	copyCmd := &CopyCommand{editor: editor}
	copyCmd.Execute()

	// Execute Cut Command
	cutCmd := &CutCommand{editor: editor}
	cutCmd.Execute()
	history.Push(cutCmd)

	// Execute Paste Command
	pasteCmd := &PasteCommand{editor: editor}
	pasteCmd.Execute()
	history.Push(pasteCmd)

	// Undo the last command
	if cmd := history.Pop(); cmd != nil {
		cmd.Undo()
	}

	// Undo another command
	if cmd := history.Pop(); cmd != nil {
		cmd.Undo()
	}
}
