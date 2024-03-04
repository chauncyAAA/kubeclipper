package generators

var CategoryTemplate = `
# <strong>{{.Name}}</strong>
`

var CommandTemplate = `
------------

# {{.MainCommand.Name}}

{{.MainCommand.Example}}

{{.MainCommand.Description}}

### Usage

` + "`" + `$ kcctl {{.MainCommand.Usage}}` + "`" + `

{{if .MainCommand.Options}}

### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- {{range $option := .MainCommand.Options}}
{{$option.Name}} | {{$option.Shorthand}} | {{$option.DefaultValue}} | {{$option.Usage}} {{end}}
{{end}}
{{$mainCommandName := .MainCommand.Name}}
{{range $sub := .SubCommands}}
------------

## <em>{{$sub.Path}}</em>

{{$sub.Example}}

{{$sub.Description}}

### Usage

` + "`" + `$ kcctl {{$mainCommandName}} {{$sub.Usage}}` + "`" + `

{{if $sub.Options}}

### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- {{range $option := $sub.Options}}
{{$option.Name}} | {{$option.Shorthand}} | {{$option.DefaultValue}} | {{$option.Usage}} {{end}}
{{end}}

{{end}}

`
