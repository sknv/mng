package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/sknv/mng/odm/document"
)

var LimitMax = 50

type (
	BaseRepository struct {
		CollectionName string
	}

	PagingParams struct {
		Limit int
		Skip  int
		Sort  []string
	}
)

func (r *BaseRepository) CollectionForDb(db *mgo.Database) *mgo.Collection {
	return db.C(r.CollectionName)
}

func (r *BaseRepository) CollectionForSession(
	session *mgo.Session,
) *mgo.Collection {
	db := session.DB("")
	return r.CollectionForDb(db)
}

func (r *BaseRepository) Find(session *mgo.Session, query bson.M) *mgo.Query {
	c := r.CollectionForSession(session)
	return c.Find(query)
}

func (r *BaseRepository) FindPage(
	session *mgo.Session, query bson.M, params PagingParams,
) *mgo.Query {
	qry := r.Find(session, query)

	// Set limit and skip params.
	limit := LimitMax
	if params.Limit > 0 && params.Limit < limit {
		limit = params.Limit // Restrict fetching limit.
	}
	qry = qry.Limit(limit)

	if params.Skip > 0 {
		qry = qry.Skip(params.Skip)
	}

	// Sort query.
	if len(params.Sort) > 0 {
		qry = qry.Sort(params.Sort...)
	}

	return qry
}

func (r *BaseRepository) Insert(session *mgo.Session, doc interface{}) error {
	col := r.CollectionForSession(session)

	// Before callbacks section.
	doBeforeInsertIfNeeded(doc)
	doBeforeSaveIfNeeded(doc)

	err := col.Insert(doc)

	// After callbacks section.
	doAfterInsertIfNeeded(doc)
	doAfterSaveIfNeeded(doc)

	return err
}

func (r *BaseRepository) Update(
	session *mgo.Session, selector interface{}, update interface{},
) error {
	col := r.CollectionForSession(session)

	// Before callbacks section.
	doBeforeUpdateIfNeeded(update)
	doBeforeSaveIfNeeded(update)

	err := col.Update(selector, update)

	// After callbacks section.
	doAfterUpdateIfNeeded(update)
	doAfterSaveIfNeeded(update)

	return err
}

func (r *BaseRepository) UpdateDoc(
	session *mgo.Session, doc document.IIdentifier,
) error {
	return r.Update(session, bson.M{"_id": doc.GetId()}, doc)
}

func (r *BaseRepository) Remove(
	session *mgo.Session, selector interface{},
) error {
	col := r.CollectionForSession(session)
	return col.Remove(selector)
}

func (r *BaseRepository) RemoveDoc(
	session *mgo.Session, doc document.IIdentifier,
) error {
	// Before callbacks section.
	doBeforeRemoveIfNeeded(doc)

	err := r.Remove(session, bson.M{"_id": doc.GetId()})

	// After callbacks section.
	doAfterRemoveIfNeeded(doc)

	return err
}
