package symfony

type SymfonyXml struct {
	Commands Commands `xml:"commands"`
}

type Commands struct {
	Command []Command `xml:"command"`
}

type Command struct {
	ID     string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Hidden int    `xml:"hidden,attr"`

	Usages      Usages    `xml:"usages"`
	Description string    `xml:"description"`
	Help        string    `xml:"help"`
	Arguments   Arguments `xml:"arguments"`
	Options     Options   `xml:"options"`
}

func (c *Command) GetOptionNames() func(string) []string {
	return func(line string) []string {
		var names []string
		for _, o := range c.Options.Option {
			AcceptValue := ""
			if o.AcceptValue {
				AcceptValue = "="
			}
			name := o.Name + AcceptValue
			names = append(names, name)
			if len(o.Shortcut) > 0 {
				names = append(names, o.Shortcut)
			}
		}
		return names
	}
}

type Usages struct {
	Usage []string `xml:"usage"`
}

type Arguments struct {
	Argument []Argument `xml:"argument"`
}

type Argument struct {
	Name       string `xml:"name,attr"`
	IsRequired bool   `xml:"is_required,attr"`
	IsArray    bool   `xml:"is_array,attr"`

	Description string   `xml:"description"`
	Defaults    Defaults `xml:"defaults"`
}

type Defaults struct {
	Default []string `xml:"default"`
}

type Options struct {
	Option []Option `xml:"option"`
}

type Option struct {
	Name            string `xml:"name,attr"`
	Shortcut        string `xml:"shortcut,attr"`
	AcceptValue     bool   `xml:"accept_value,attr"`
	IsValueRequired bool   `xml:"is_value_required,attr"`
	IsMultiple      bool   `xml:"is_multiple,attr"`

	Description string   `xml:"description"`
	Defaults    Defaults `xml:"defaults"`
}
