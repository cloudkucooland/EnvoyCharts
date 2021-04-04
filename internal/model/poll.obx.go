// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package model

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type entry_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var EntryBinding = entry_EntityInfo{
	Entity: objectbox.Entity{
		Id: 1,
	},
	Uid: 2970823394468775929,
}

// Entry_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Entry_ = struct {
	Id           *objectbox.PropertyInt64
	Date         *objectbox.PropertyInt64
	ProductionW  *objectbox.PropertyFloat64
	ConsumptionW *objectbox.PropertyFloat64
	NetW         *objectbox.PropertyFloat64
}{
	Id: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &EntryBinding.Entity,
		},
	},
	Date: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &EntryBinding.Entity,
		},
	},
	ProductionW: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &EntryBinding.Entity,
		},
	},
	ConsumptionW: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &EntryBinding.Entity,
		},
	},
	NetW: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &EntryBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (entry_EntityInfo) GeneratorVersion() int {
	return 5
}

// AddToModel is called by ObjectBox during model build
func (entry_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Entry", 1, 2970823394468775929)
	model.Property("Id", 6, 1, 1732099957991171956)
	model.PropertyFlags(1)
	model.Property("Date", 6, 2, 2580885015314113201)
	model.Property("ProductionW", 8, 3, 2190537801956453137)
	model.Property("ConsumptionW", 8, 4, 2527778538083053632)
	model.Property("NetW", 8, 5, 2946360827411257486)
	model.EntityLastPropertyId(5, 2946360827411257486)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (entry_EntityInfo) GetId(object interface{}) (uint64, error) {
	return uint64(object.(*Entry).Id), nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (entry_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Entry).Id = int64(id)
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (entry_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (entry_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Entry)

	// build the FlatBuffers object
	fbb.StartObject(5)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, obj.Date)
	fbutils.SetFloat64Slot(fbb, 2, obj.ProductionW)
	fbutils.SetFloat64Slot(fbb, 3, obj.ConsumptionW)
	fbutils.SetFloat64Slot(fbb, 4, obj.NetW)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (entry_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Entry' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetInt64Slot(4, 0)

	return &Entry{
		Id:           propId,
		Date:         fbutils.GetInt64Slot(table, 6),
		ProductionW:  fbutils.GetFloat64Slot(table, 8),
		ConsumptionW: fbutils.GetFloat64Slot(table, 10),
		NetW:         fbutils.GetFloat64Slot(table, 12),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (entry_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Entry, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (entry_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Entry), nil)
	}
	return append(slice.([]*Entry), object.(*Entry))
}

// Box provides CRUD access to Entry objects
type EntryBox struct {
	*objectbox.Box
}

// BoxForEntry opens a box of Entry objects
func BoxForEntry(ob *objectbox.ObjectBox) *EntryBox {
	return &EntryBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Entry.Id property on the passed object will be assigned the new ID as well.
func (box *EntryBox) Put(object *Entry) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Entry.Id property on the passed object will be assigned the new ID as well.
func (box *EntryBox) Insert(object *Entry) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *EntryBox) Update(object *Entry) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *EntryBox) PutAsync(object *Entry) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Entry.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Entry.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *EntryBox) PutMany(objects []*Entry) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *EntryBox) Get(id uint64) (*Entry, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Entry), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *EntryBox) GetMany(ids ...uint64) ([]*Entry, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Entry), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *EntryBox) GetManyExisting(ids ...uint64) ([]*Entry, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Entry), nil
}

// GetAll reads all stored objects
func (box *EntryBox) GetAll() ([]*Entry, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Entry), nil
}

// Remove deletes a single object
func (box *EntryBox) Remove(object *Entry) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *EntryBox) RemoveMany(objects ...*Entry) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = uint64(object.Id)
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Entry_ struct to create conditions.
// Keep the *EntryQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *EntryBox) Query(conditions ...objectbox.Condition) *EntryQuery {
	return &EntryQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Entry_ struct to create conditions.
// Keep the *EntryQuery if you intend to execute the query multiple times.
func (box *EntryBox) QueryOrError(conditions ...objectbox.Condition) (*EntryQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &EntryQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See EntryAsyncBox for more information.
func (box *EntryBox) Async() *EntryAsyncBox {
	return &EntryAsyncBox{AsyncBox: box.Box.Async()}
}

// EntryAsyncBox provides asynchronous operations on Entry objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type EntryAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForEntry creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use EntryBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForEntry(ob *objectbox.ObjectBox, timeoutMs uint64) *EntryAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 1, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 1: %s" + err.Error())
	}
	return &EntryAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *EntryAsyncBox) Put(object *Entry) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *EntryAsyncBox) Insert(object *Entry) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *EntryAsyncBox) Update(object *Entry) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *EntryAsyncBox) Remove(object *Entry) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Entry which Id is either 42 or 47:
// 		box.Query(Entry_.Id.In(42, 47)).Find()
type EntryQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *EntryQuery) Find() ([]*Entry, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Entry), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *EntryQuery) Offset(offset uint64) *EntryQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *EntryQuery) Limit(limit uint64) *EntryQuery {
	query.Query.Limit(limit)
	return query
}

type daily_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var DailyBinding = daily_EntityInfo{
	Entity: objectbox.Entity{
		Id: 2,
	},
	Uid: 6434397660819005171,
}

// Daily_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Daily_ = struct {
	Date           *objectbox.PropertyInt64
	ConsumptionkWh *objectbox.PropertyFloat64
	DID            *objectbox.PropertyInt64
	ProductionkWh  *objectbox.PropertyFloat64
}{
	Date: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &DailyBinding.Entity,
		},
	},
	ConsumptionkWh: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &DailyBinding.Entity,
		},
	},
	DID: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &DailyBinding.Entity,
		},
	},
	ProductionkWh: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &DailyBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (daily_EntityInfo) GeneratorVersion() int {
	return 5
}

// AddToModel is called by ObjectBox during model build
func (daily_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Daily", 2, 6434397660819005171)
	model.Property("Date", 10, 2, 5037411518101750911)
	model.Property("ConsumptionkWh", 8, 4, 3504309473206441362)
	model.Property("DID", 6, 5, 1048918584752691851)
	model.PropertyFlags(129)
	model.Property("ProductionkWh", 8, 6, 3218371933083291281)
	model.EntityLastPropertyId(6, 3218371933083291281)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (daily_EntityInfo) GetId(object interface{}) (uint64, error) {
	return uint64(object.(*Daily).DID), nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (daily_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Daily).DID = int64(id)
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (daily_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (daily_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Daily)
	var propDate int64
	{
		var err error
		propDate, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.Date)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Daily.Date: " + err.Error())
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(6)
	fbutils.SetUint64Slot(fbb, 4, id)
	fbutils.SetInt64Slot(fbb, 1, propDate)
	fbutils.SetFloat64Slot(fbb, 5, obj.ProductionkWh)
	fbutils.SetFloat64Slot(fbb, 3, obj.ConsumptionkWh)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (daily_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Daily' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propDID = table.GetInt64Slot(12, 0)

	propDate, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 6))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Daily.Date: " + err.Error())
	}

	return &Daily{
		DID:            propDID,
		Date:           propDate,
		ProductionkWh:  fbutils.GetFloat64Slot(table, 14),
		ConsumptionkWh: fbutils.GetFloat64Slot(table, 10),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (daily_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Daily, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (daily_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Daily), nil)
	}
	return append(slice.([]*Daily), object.(*Daily))
}

// Box provides CRUD access to Daily objects
type DailyBox struct {
	*objectbox.Box
}

// BoxForDaily opens a box of Daily objects
func BoxForDaily(ob *objectbox.ObjectBox) *DailyBox {
	return &DailyBox{
		Box: ob.InternalBox(2),
	}
}

// Put synchronously inserts/updates a single object.
// In case the DID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Daily.DID property on the passed object will be assigned the new ID as well.
func (box *DailyBox) Put(object *Daily) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the DID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Daily.DID property on the passed object will be assigned the new ID as well.
func (box *DailyBox) Insert(object *Daily) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *DailyBox) Update(object *Daily) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *DailyBox) PutAsync(object *Daily) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case DIDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Daily.DID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Daily.DID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *DailyBox) PutMany(objects []*Daily) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *DailyBox) Get(id uint64) (*Daily, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Daily), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *DailyBox) GetMany(ids ...uint64) ([]*Daily, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Daily), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *DailyBox) GetManyExisting(ids ...uint64) ([]*Daily, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Daily), nil
}

// GetAll reads all stored objects
func (box *DailyBox) GetAll() ([]*Daily, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Daily), nil
}

// Remove deletes a single object
func (box *DailyBox) Remove(object *Daily) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *DailyBox) RemoveMany(objects ...*Daily) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = uint64(object.DID)
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Daily_ struct to create conditions.
// Keep the *DailyQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *DailyBox) Query(conditions ...objectbox.Condition) *DailyQuery {
	return &DailyQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Daily_ struct to create conditions.
// Keep the *DailyQuery if you intend to execute the query multiple times.
func (box *DailyBox) QueryOrError(conditions ...objectbox.Condition) (*DailyQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &DailyQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See DailyAsyncBox for more information.
func (box *DailyBox) Async() *DailyAsyncBox {
	return &DailyAsyncBox{AsyncBox: box.Box.Async()}
}

// DailyAsyncBox provides asynchronous operations on Daily objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type DailyAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForDaily creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use DailyBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForDaily(ob *objectbox.ObjectBox, timeoutMs uint64) *DailyAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 2, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 2: %s" + err.Error())
	}
	return &DailyAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the DID property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *DailyAsyncBox) Put(object *Daily) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The DID property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *DailyAsyncBox) Insert(object *Daily) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *DailyAsyncBox) Update(object *Daily) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *DailyAsyncBox) Remove(object *Daily) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Daily which DID is either 42 or 47:
// 		box.Query(Daily_.DID.In(42, 47)).Find()
type DailyQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *DailyQuery) Find() ([]*Daily, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Daily), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *DailyQuery) Offset(offset uint64) *DailyQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *DailyQuery) Limit(limit uint64) *DailyQuery {
	query.Query.Limit(limit)
	return query
}
