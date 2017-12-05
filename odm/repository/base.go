package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/sknv/mng/odm/document"
)

var MaxFetchLimit = 50

type Base struct {
	DbName         string
	CollectionName string
}

func (r *Base) DB(session *mgo.Session) *mgo.Database {
	return session.DB(r.DbName)
}

func (r *Base) CollectionForDb(db *mgo.Database) *mgo.Collection {
	return db.C(r.CollectionName)
}

func (r *Base) CollectionForSession(
	session *mgo.Session,
) *mgo.Collection {
	db := r.DB(session)
	return r.CollectionForDb(db)
}

func (r *Base) Find(session *mgo.Session, query bson.M) *mgo.Query {
	c := r.CollectionForSession(session)
	return c.Find(query)
}

func (r *Base) FindPage(
	session *mgo.Session, query bson.M, limit, skip int, result interface{},
) error {
	q := r.Find(session, query)

	// Set limit and skip params.
	if limit > 0 {
		// Restrict fetching limit.
		if limit > MaxFetchLimit {
			limit = MaxFetchLimit
		}
		q.Limit(limit)
	}
	if skip > 0 {
		q.Skip(skip)
	}

	return q.All(result)
}

func (r *Base) FindAll(
	session *mgo.Session, query bson.M, result interface{},
) error {
	return r.FindPage(session, query, 0, 0, result)
}

func (r *Base) Page(
	session *mgo.Session, limit, skip int, result interface{},
) error {
	return r.FindPage(session, bson.M{}, limit, skip, result)
}

func (r *Base) All(session *mgo.Session, result interface{}) error {
	return r.FindAll(session, bson.M{}, result)
}

func (r *Base) FindOne(
	session *mgo.Session, query bson.M, result interface{},
) error {
	q := r.Find(session, query)
	return q.One(result)
}

func (r *Base) FindOneById(
	session *mgo.Session, id interface{}, result interface{},
) error {
	return r.FindOne(session, bson.M{"_id": id}, result)
}

func (r *Base) Insert(session *mgo.Session, doc interface{}) error {
	c := r.CollectionForSession(session)

	// Before callbacks section.
	doBeforeInsertIfNeeded(doc)
	doBeforeSaveIfNeeded(doc)

	err := c.Insert(doc)

	// After callbacks section.
	doAfterInsertIfNeeded(doc)
	doAfterSaveIfNeeded(doc)

	return err
}

func (r *Base) Update(
	session *mgo.Session, selector interface{}, update interface{},
) error {
	c := r.CollectionForSession(session)

	// Before callbacks section.
	doBeforeUpdateIfNeeded(update)
	doBeforeSaveIfNeeded(update)

	err := c.Update(selector, update)

	// After callbacks section.
	doAfterUpdateIfNeeded(update)
	doAfterSaveIfNeeded(update)

	return err
}

func (r *Base) UpdateById(
	session *mgo.Session, id interface{}, update interface{},
) error {
	return r.Update(session, bson.M{"_id": id}, update)
}

func (r *Base) UpdateDoc(session *mgo.Session, doc document.IIdentifier) error {
	return r.UpdateById(session, doc.GetId(), doc)
}

func (r *Base) Remove(session *mgo.Session, selector interface{}) error {
	c := r.CollectionForSession(session)
	return c.Remove(selector)
}

func (r *Base) RemoveById(session *mgo.Session, id interface{}) error {
	return r.Remove(session, bson.M{"_id": id})
}

func (r *Base) RemoveDoc(session *mgo.Session, doc document.IIdentifier) error {
	// Before callbacks section.
	doBeforeRemoveIfNeeded(doc)

	err := r.RemoveById(session, doc.GetId())

	// After callbacks section.
	doAfterRemoveIfNeeded(doc)

	return err
}
