package local

import (
	"github.com/globocom/commandmocker"
	"github.com/globocom/tsuru/provision"
	"github.com/globocom/tsuru/testing"
	. "launchpad.net/gocheck"
)

func (s *S) TestProvisionerProvision(c *C) {
	tmpdir, err := commandmocker.Add("sudo", "$*")
	c.Assert(err, IsNil)
	defer commandmocker.Remove(tmpdir)
	var p LocalProvisioner
	app := testing.NewFakeApp("myapp", "python", 0)
	c.Assert(p.Provision(app), IsNil)
	c.Assert(commandmocker.Ran(tmpdir), Equals, true)
	expected := "lxc-create -t ubuntu -n myapp"
	expected += "lxc-start --daemon -n myapp"
	c.Assert(commandmocker.Output(tmpdir), Equals, expected)
}

func (s *S) TestProvisionerDestroy(c *C) {
	tmpdir, err := commandmocker.Add("sudo", "$*")
	c.Assert(err, IsNil)
	defer commandmocker.Remove(tmpdir)
	var p LocalProvisioner
	app := testing.NewFakeApp("myapp", "python", 0)
	c.Assert(p.Destroy(app), IsNil)
	c.Assert(commandmocker.Ran(tmpdir), Equals, true)
	expected := "lxc-stop -n myapp"
	expected += "lxc-destroy -n myapp"
	c.Assert(commandmocker.Output(tmpdir), Equals, expected)
}

func (s *S) TestProvisionerAddr(c *C) {
	var p LocalProvisioner
	app := testing.NewFakeApp("myapp", "python", 1)
	addr, err := p.Addr(app)
	c.Assert(err, IsNil)
	c.Assert(addr, Equals, app.ProvisionUnits()[0].GetIp())
}

func (s *S) TestProvisionerAddUnits(c *C) {
	var p LocalProvisioner
	app := testing.NewFakeApp("myapp", "python", 0)
	units, err := p.AddUnits(app, 2)
	c.Assert(err, IsNil)
	c.Assert(units, DeepEquals, []provision.Unit{})
}

func (s *S) TestProvisionerRemoveUnit(c *C) {
	var p LocalProvisioner
	app := testing.NewFakeApp("myapp", "python", 0)
	err := p.RemoveUnit(app, "")
	c.Assert(err, IsNil)
}
