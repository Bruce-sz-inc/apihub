package account

import (
	"github.com/albertoleal/backstage/errors"
	. "gopkg.in/check.v1"
)

var (
	team *Team
	owner *User
)

func (s *S) SetUpTest(c *C) {
	team = &Team{Name: "Team"}
	owner = &User{Name: "Alice", Username: "alice"}
}

func (s *S) TestCreateTeam(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	c.Assert(err, IsNil)
}

func (s *S) TestCreateTeamWhenNameAlreadyExists(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	c.Assert(err, IsNil)

	team = &Team{Name: "Team"}
	err = team.Save(owner)
	c.Assert(err, NotNil)
	e := err.(*errors.ValidationError)
	message := "Someone already has that team name. Could you try another?"
	c.Assert(e.Message, Equals, message)
}

func (s *S) TestDeleteTeam(c *C) {
	teamName := team.Name
	team.Save(owner)
	g, _ := FindTeamByName(team.Name)
	c.Assert(len(g.Users), Equals, 1)
	team.Delete()
	_, err := FindTeamByName(teamName)
	c.Assert(err, Not(IsNil))
}

func (s *S) TestAddUsersWithInvalidUser(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	g, _ := FindTeamByName("Team")

	err = g.AddUsers([]string{"alice", "bob"})
	c.Assert(err, IsNil)

	g, _ = FindTeamByName("Team")
	c.Assert(len(g.Users), Equals, 1)
}

func (s *S) TestAddUsersWithValidUser(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	g, _ := FindTeamByName("Team")

	bob := &User{Name: "Bob", Email: "bob@bar.com", Username: "bob", Password: "123456"}
	bob.Save()
	defer bob.Delete()
	err = g.AddUsers([]string{"alice", "bob"})
	c.Assert(err, IsNil)

	g, _ = FindTeamByName("Team")
	c.Assert(len(g.Users), Equals, 2)
}

func (s *S) TestAddUsersWithSameUsername(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	g, _ := FindTeamByName("Team")
	c.Assert(len(g.Users), Equals, 1)

	err = g.AddUsers([]string{"alice", "alice"})
	c.Assert(err, IsNil)

	g, _ = FindTeamByName("Team")
	c.Assert(len(g.Users), Equals, 1)
}

func (s *S) TestRemoveUsers(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	g, _ := FindTeamByName("Team")
	err = g.AddUsers([]string{"alice", "bob"})
	err = g.RemoveUsers([]string{"bob"})
	c.Assert(err, IsNil)
	g, _ = FindTeamByName("Team")
	c.Assert(len(g.Users), Equals, 1)
	c.Assert(g.Users[0], Equals, "alice")
}

func (s *S) TestRemoveUsersWithNonExistingUser(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	g, _ := FindTeamByName("Team")
	err = g.RemoveUsers([]string{"bob"})
	c.Assert(err, IsNil)
}

func (s *S) TestRemoveUsersWhenTheUserIsOwner(c *C) {
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	mary := &User{Name: "Mary", Email: "mary@bar.com", Username: "mary", Password: "123456"}
	mary.Save()
	defer mary.Delete()
	team.AddUsers([]string{"mary", "bob"})

	err = team.RemoveUsers([]string{owner.Username, "bob"})
	c.Assert(err, Not(IsNil))
	e := err.(*errors.ValidationError)
	c.Assert(e.Message, Equals, "It is not possible to remove the owner from the team.")

	g, _ := FindTeamByName("Team")
	c.Assert(len(g.Users), Equals, 2)
	c.Assert(g.Users[0], Equals, owner.Username)
}

func (s *S) TestDeleteTeamByName(c *C) {
	err := team.Save(owner)
	c.Assert(err, IsNil)
	err = DeleteTeamByName("Team")
	c.Assert(err, IsNil)
}

func (s *S) TestDeleteTeamByNameWithInvalidName(c *C) {
	err := DeleteTeamByName("Non Existing Team")
	c.Assert(err, NotNil)
	e := err.(*errors.ValidationError)
	message := "Team not found."
	c.Assert(e.Message, Equals, message)
}

func (s *S) TestFindTeamByName(c *C) {
	owner := &User{Name: "Alice", Email: "alice@bar.com", Username: "alice", Password: "123456"}
	err := team.Save(owner)

	defer DeleteTeamByName("Team")
	c.Assert(err, IsNil)

	g, _ := FindTeamByName("Team")
	c.Assert(g.Name, Equals, "Team")
}

func (s *S) TestFindTeamByNameWithInvalidName(c *C) {
	_, err := FindTeamByName("Non Existing Team")
	c.Assert(err, NotNil)
	e := err.(*errors.ValidationError)
	message := "Team not found."
	c.Assert(e.Message, Equals, message)
}

func (s *S) TestFindTeamById(c *C) {
	owner := &User{Name: "Alice", Email: "alice@bar.com", Username: "alice", Password: "123456"}
	err := team.Save(owner)
	defer DeleteTeamByName("Team")
	c.Assert(err, IsNil)
	team, _ := FindTeamByName("Team")
	g, _ := FindTeamById(team.Id.Hex())
	c.Assert(g.Name, Equals, "Team")
}

func (s *S) TestFindTeamByIdWithInvalidId(c *C) {
	_, err := FindTeamById("123")
	c.Assert(err, NotNil)
	e := err.(*errors.ValidationError)
	message := "Team not found."
	c.Assert(e.Message, Equals, message)
}

func (s *S) TestGetTeamUsers(c *C) {
	alice := &User{Name: "Alice", Email: "alice@bar.com", Username: "alice", Password: "123456"}
	defer alice.Delete()
	alice.Save()
	bob := &User{Name: "Bob", Email: "bob@bar.com", Username: "bob", Password: "123456"}
	defer bob.Delete()
	bob.Save()

	team.Save(alice)
	team.AddUsers([]string{"bob"})
	defer DeleteTeamByName("Team")

	g, _ := FindTeamByName("Team")
	users, _ := g.GetTeamUsers()
	c.Assert(users[0].Username, Equals, alice.Username)
	c.Assert(users[1].Username, Equals, bob.Username)
}

func (s *S) TestContainsUser(c *C) {
	bob := &User{Name: "Bob", Email: "bob@bar.com", Username: "bob", Password: "123456"}
	defer bob.Delete()
	bob.Save()

	team.Save(owner)
	defer DeleteTeamByName("Team")
	g, _ := FindTeamByName("Team")
	_, ok := g.ContainsUser(owner)
	c.Assert(ok, Equals, true)
	_, ok = g.ContainsUser(bob)
	c.Assert(ok, Equals, false)
}