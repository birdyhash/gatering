package gatering

/*
package gatering contains the core execution chamber.

a chamber is responsible for moving a payload through
a sequence of gates.

each gate receives the output of the previous gate and
returns a transformed version of the payload.

the chamber also records the route taken by the payload,
allowing the caller to inspect the execution path after
processing is complete.
*/
type Gate struct {
	Name string
	Run  func([]byte) []byte
}

type Chamber struct {
	route []Gate
	logs  []string
}

/*
new creates an empty chamber.

the returned chamber does not contain any gates and
must be configured before use.
*/
func New() *Chamber {
	return &Chamber{}
}

/*
add registers a gate inside the chamber.

gates are executed in the same order they are added.
*/
func (c *Chamber) Add(g Gate) {
	c.route = append(c.route, g)
}

/*
pass sends a payload through every configured gate.

the final transformed payload is returned to the caller.
*/
func (c *Chamber) Pass(data []byte) []byte {

	current := data

	c.logs = nil

	for _, gate := range c.route {

		c.logs = append(c.logs, gate.Name)

		current = gate.Run(current)
	}

	return current
}

/*
route returns the list of gates crossed during
the most recent execution.
*/
func (c *Chamber) Route() []string {
	return c.logs
}
