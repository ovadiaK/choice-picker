# choice-picker

Example program to present generic buttons representing choices. Every choice consists of two parts: A text displayed to the user and a function run when user pressed the corresponding button.

### state
The program remembers its state, here a history of pressed buttons and the current choices.
```go
type State struct {
	History history
	Current []Choice
}
```
The current choices are build in to the template and visible to the user.
```go
{{range  $i, $e := .Current}}
        <form action="/" method="get">
            <input type="number" value="{{$i}}" name="choice" readonly hidden>
            <button type="submit">{{$e.Text}}</button>
        </form>
    {{end}}
```
### choices
Choices consist of two parts: one displayed to the user, one function to be run by the backend.
```go
type Choice struct {
	Text   string
	result func(*history)
}
```
When the user's choice was identified, the correct function is fired, doing whatever it does.

```go
func (s *State) getChoice(r *http.Request) error {
	n, err := strconv.Atoi(r.FormValue("choice"))
	if err != nil {
		return err
	}
	s.Current[n].result(&s.History)
	return nil
}
```
In this example it's just adding letters...
```go
result: func(h *history) {
			*h += history(possibilities[n])
		},
```
...but the possibilities are infinte!
