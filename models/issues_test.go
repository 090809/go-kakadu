// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testIssues(t *testing.T) {
	t.Parallel()

	query := Issues()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIssuesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIssuesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Issues().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIssuesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IssueSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIssuesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IssueExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Issue exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IssueExists to return true, but got false.")
	}
}

func testIssuesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	issueFound, err := FindIssue(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if issueFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIssuesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Issues().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIssuesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Issues().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIssuesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	issueOne := &Issue{}
	issueTwo := &Issue{}
	if err = randomize.Struct(seed, issueOne, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}
	if err = randomize.Struct(seed, issueTwo, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = issueOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = issueTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Issues().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIssuesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	issueOne := &Issue{}
	issueTwo := &Issue{}
	if err = randomize.Struct(seed, issueOne, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}
	if err = randomize.Struct(seed, issueTwo, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = issueOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = issueTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func issueBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func issueAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
	*o = Issue{}
	return nil
}

func testIssuesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Issue{}
	o := &Issue{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, issueDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Issue object: %s", err)
	}

	AddIssueHook(boil.BeforeInsertHook, issueBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	issueBeforeInsertHooks = []IssueHook{}

	AddIssueHook(boil.AfterInsertHook, issueAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	issueAfterInsertHooks = []IssueHook{}

	AddIssueHook(boil.AfterSelectHook, issueAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	issueAfterSelectHooks = []IssueHook{}

	AddIssueHook(boil.BeforeUpdateHook, issueBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	issueBeforeUpdateHooks = []IssueHook{}

	AddIssueHook(boil.AfterUpdateHook, issueAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	issueAfterUpdateHooks = []IssueHook{}

	AddIssueHook(boil.BeforeDeleteHook, issueBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	issueBeforeDeleteHooks = []IssueHook{}

	AddIssueHook(boil.AfterDeleteHook, issueAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	issueAfterDeleteHooks = []IssueHook{}

	AddIssueHook(boil.BeforeUpsertHook, issueBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	issueBeforeUpsertHooks = []IssueHook{}

	AddIssueHook(boil.AfterUpsertHook, issueAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	issueAfterUpsertHooks = []IssueHook{}
}

func testIssuesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIssuesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(issueColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIssueToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into `issue_user` (`issue_id`, `user_id`) values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into `issue_user` (`issue_id`, `user_id`) values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := IssueSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*Issue)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testIssueToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Issues[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Issues[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testIssueToManySetOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUsers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUsers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Issues) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Issues) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Issues[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Issues[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Users[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Users[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testIssueToManyRemoveOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUsers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUsers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Issues) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Issues) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Issues[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Issues[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Users) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Users[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Users[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testIssueToOneUserUsingPerformer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Issue
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.PerformerID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Performer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := IssueSlice{&local}
	if err = local.L.LoadPerformer(ctx, tx, false, (*[]*Issue)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Performer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Performer = nil
	if err = local.L.LoadPerformer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Performer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIssueToOneUserUsingCreator(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Issue
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CreatorID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Creator().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := IssueSlice{&local}
	if err = local.L.LoadCreator(ctx, tx, false, (*[]*Issue)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Creator == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Creator = nil
	if err = local.L.LoadCreator(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Creator == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIssueToOneProjectUsingProject(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Issue
	var foreign Project

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ProjectID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Project().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := IssueSlice{&local}
	if err = local.L.LoadProject(ctx, tx, false, (*[]*Issue)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Project == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Project = nil
	if err = local.L.LoadProject(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Project == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIssueToOneSetOpUserUsingPerformer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetPerformer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Performer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PerformerIssues[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.PerformerID, x.ID) {
			t.Error("foreign key was wrong value", a.PerformerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PerformerID))
		reflect.Indirect(reflect.ValueOf(&a.PerformerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.PerformerID, x.ID) {
			t.Error("foreign key was wrong value", a.PerformerID, x.ID)
		}
	}
}

func testIssueToOneRemoveOpUserUsingPerformer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPerformer(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePerformer(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Performer().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Performer != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.PerformerID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PerformerIssues) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIssueToOneSetOpUserUsingCreator(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetCreator(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Creator != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CreatorIssues[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CreatorID, x.ID) {
			t.Error("foreign key was wrong value", a.CreatorID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CreatorID))
		reflect.Indirect(reflect.ValueOf(&a.CreatorID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CreatorID, x.ID) {
			t.Error("foreign key was wrong value", a.CreatorID, x.ID)
		}
	}
}

func testIssueToOneRemoveOpUserUsingCreator(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCreator(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCreator(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Creator().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Creator != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.CreatorID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CreatorIssues) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIssueToOneSetOpProjectUsingProject(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Project{&b, &c} {
		err = a.SetProject(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Project != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Issues[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ProjectID, x.ID) {
			t.Error("foreign key was wrong value", a.ProjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProjectID))
		reflect.Indirect(reflect.ValueOf(&a.ProjectID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ProjectID, x.ID) {
			t.Error("foreign key was wrong value", a.ProjectID, x.ID)
		}
	}
}

func testIssueToOneRemoveOpProjectUsingProject(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Issue
	var b Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetProject(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveProject(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Project().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Project != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ProjectID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Issues) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIssuesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIssuesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IssueSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIssuesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Issues().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	issueDBTypes = map[string]string{`ID`: `int`, `Name`: `tinytext`, `Description`: `text`, `PerformerID`: `int`, `CreatorID`: `int`, `ProjectID`: `int`, `Status`: `enum('NOT_STARTED','IN_PROGRESS','DONE')`, `DeadlineAt`: `datetime`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_            = bytes.MinRead
)

func testIssuesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(issuePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(issueAllColumns) == len(issuePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, issueDBTypes, true, issuePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIssuesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(issueAllColumns) == len(issuePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Issue{}
	if err = randomize.Struct(seed, o, issueDBTypes, true, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, issueDBTypes, true, issuePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(issueAllColumns, issuePrimaryKeyColumns) {
		fields = issueAllColumns
	} else {
		fields = strmangle.SetComplement(
			issueAllColumns,
			issuePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := IssueSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIssuesUpsert(t *testing.T) {
	t.Parallel()

	if len(issueAllColumns) == len(issuePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLIssueUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Issue{}
	if err = randomize.Struct(seed, &o, issueDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Issue: %s", err)
	}

	count, err := Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, issueDBTypes, false, issuePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Issue: %s", err)
	}

	count, err = Issues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}