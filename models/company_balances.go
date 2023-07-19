// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CompanyBalance is an object representing the database table.
type CompanyBalance struct {
	ID        int64           `boil:"id" json:"id" toml:"id" yaml:"id"`
	CompanyID int64           `boil:"company_id" json:"company_id" toml:"company_id" yaml:"company_id"`
	Balance   decimal.Decimal `boil:"balance" json:"balance" toml:"balance" yaml:"balance"`

	R *companyBalanceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L companyBalanceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompanyBalanceColumns = struct {
	ID        string
	CompanyID string
	Balance   string
}{
	ID:        "id",
	CompanyID: "company_id",
	Balance:   "balance",
}

var CompanyBalanceTableColumns = struct {
	ID        string
	CompanyID string
	Balance   string
}{
	ID:        "company_balances.id",
	CompanyID: "company_balances.company_id",
	Balance:   "company_balances.balance",
}

// Generated where

var CompanyBalanceWhere = struct {
	ID        whereHelperint64
	CompanyID whereHelperint64
	Balance   whereHelperdecimal_Decimal
}{
	ID:        whereHelperint64{field: "`company_balances`.`id`"},
	CompanyID: whereHelperint64{field: "`company_balances`.`company_id`"},
	Balance:   whereHelperdecimal_Decimal{field: "`company_balances`.`balance`"},
}

// CompanyBalanceRels is where relationship names are stored.
var CompanyBalanceRels = struct {
}{}

// companyBalanceR is where relationships are stored.
type companyBalanceR struct {
}

// NewStruct creates a new relationship struct
func (*companyBalanceR) NewStruct() *companyBalanceR {
	return &companyBalanceR{}
}

// companyBalanceL is where Load methods for each relationship are stored.
type companyBalanceL struct{}

var (
	companyBalanceAllColumns            = []string{"id", "company_id", "balance"}
	companyBalanceColumnsWithoutDefault = []string{"company_id"}
	companyBalanceColumnsWithDefault    = []string{"id", "balance"}
	companyBalancePrimaryKeyColumns     = []string{"id"}
	companyBalanceGeneratedColumns      = []string{}
)

type (
	// CompanyBalanceSlice is an alias for a slice of pointers to CompanyBalance.
	// This should almost always be used instead of []CompanyBalance.
	CompanyBalanceSlice []*CompanyBalance
	// CompanyBalanceHook is the signature for custom CompanyBalance hook methods
	CompanyBalanceHook func(context.Context, boil.ContextExecutor, *CompanyBalance) error

	companyBalanceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	companyBalanceType                 = reflect.TypeOf(&CompanyBalance{})
	companyBalanceMapping              = queries.MakeStructMapping(companyBalanceType)
	companyBalancePrimaryKeyMapping, _ = queries.BindMapping(companyBalanceType, companyBalanceMapping, companyBalancePrimaryKeyColumns)
	companyBalanceInsertCacheMut       sync.RWMutex
	companyBalanceInsertCache          = make(map[string]insertCache)
	companyBalanceUpdateCacheMut       sync.RWMutex
	companyBalanceUpdateCache          = make(map[string]updateCache)
	companyBalanceUpsertCacheMut       sync.RWMutex
	companyBalanceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var companyBalanceAfterSelectHooks []CompanyBalanceHook

var companyBalanceBeforeInsertHooks []CompanyBalanceHook
var companyBalanceAfterInsertHooks []CompanyBalanceHook

var companyBalanceBeforeUpdateHooks []CompanyBalanceHook
var companyBalanceAfterUpdateHooks []CompanyBalanceHook

var companyBalanceBeforeDeleteHooks []CompanyBalanceHook
var companyBalanceAfterDeleteHooks []CompanyBalanceHook

var companyBalanceBeforeUpsertHooks []CompanyBalanceHook
var companyBalanceAfterUpsertHooks []CompanyBalanceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CompanyBalance) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CompanyBalance) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CompanyBalance) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CompanyBalance) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CompanyBalance) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CompanyBalance) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CompanyBalance) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CompanyBalance) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CompanyBalance) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBalanceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompanyBalanceHook registers your hook function for all future operations.
func AddCompanyBalanceHook(hookPoint boil.HookPoint, companyBalanceHook CompanyBalanceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		companyBalanceAfterSelectHooks = append(companyBalanceAfterSelectHooks, companyBalanceHook)
	case boil.BeforeInsertHook:
		companyBalanceBeforeInsertHooks = append(companyBalanceBeforeInsertHooks, companyBalanceHook)
	case boil.AfterInsertHook:
		companyBalanceAfterInsertHooks = append(companyBalanceAfterInsertHooks, companyBalanceHook)
	case boil.BeforeUpdateHook:
		companyBalanceBeforeUpdateHooks = append(companyBalanceBeforeUpdateHooks, companyBalanceHook)
	case boil.AfterUpdateHook:
		companyBalanceAfterUpdateHooks = append(companyBalanceAfterUpdateHooks, companyBalanceHook)
	case boil.BeforeDeleteHook:
		companyBalanceBeforeDeleteHooks = append(companyBalanceBeforeDeleteHooks, companyBalanceHook)
	case boil.AfterDeleteHook:
		companyBalanceAfterDeleteHooks = append(companyBalanceAfterDeleteHooks, companyBalanceHook)
	case boil.BeforeUpsertHook:
		companyBalanceBeforeUpsertHooks = append(companyBalanceBeforeUpsertHooks, companyBalanceHook)
	case boil.AfterUpsertHook:
		companyBalanceAfterUpsertHooks = append(companyBalanceAfterUpsertHooks, companyBalanceHook)
	}
}

// One returns a single companyBalance record from the query.
func (q companyBalanceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CompanyBalance, error) {
	o := &CompanyBalance{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for company_balances")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CompanyBalance records from the query.
func (q companyBalanceQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompanyBalanceSlice, error) {
	var o []*CompanyBalance

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CompanyBalance slice")
	}

	if len(companyBalanceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CompanyBalance records in the query.
func (q companyBalanceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count company_balances rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q companyBalanceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if company_balances exists")
	}

	return count > 0, nil
}

// CompanyBalances retrieves all the records using an executor.
func CompanyBalances(mods ...qm.QueryMod) companyBalanceQuery {
	mods = append(mods, qm.From("`company_balances`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`company_balances`.*"})
	}

	return companyBalanceQuery{q}
}

// FindCompanyBalance retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompanyBalance(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CompanyBalance, error) {
	companyBalanceObj := &CompanyBalance{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `company_balances` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, companyBalanceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from company_balances")
	}

	if err = companyBalanceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return companyBalanceObj, err
	}

	return companyBalanceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompanyBalance) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company_balances provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyBalanceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	companyBalanceInsertCacheMut.RLock()
	cache, cached := companyBalanceInsertCache[key]
	companyBalanceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			companyBalanceAllColumns,
			companyBalanceColumnsWithDefault,
			companyBalanceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(companyBalanceType, companyBalanceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(companyBalanceType, companyBalanceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `company_balances` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `company_balances` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `company_balances` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, companyBalancePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into company_balances")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == companyBalanceMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for company_balances")
	}

CacheNoHooks:
	if !cached {
		companyBalanceInsertCacheMut.Lock()
		companyBalanceInsertCache[key] = cache
		companyBalanceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CompanyBalance.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompanyBalance) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	companyBalanceUpdateCacheMut.RLock()
	cache, cached := companyBalanceUpdateCache[key]
	companyBalanceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			companyBalanceAllColumns,
			companyBalancePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update company_balances, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `company_balances` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, companyBalancePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(companyBalanceType, companyBalanceMapping, append(wl, companyBalancePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update company_balances row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for company_balances")
	}

	if !cached {
		companyBalanceUpdateCacheMut.Lock()
		companyBalanceUpdateCache[key] = cache
		companyBalanceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q companyBalanceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for company_balances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for company_balances")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompanyBalanceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyBalancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `company_balances` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, companyBalancePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in companyBalance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all companyBalance")
	}
	return rowsAff, nil
}

var mySQLCompanyBalanceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompanyBalance) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company_balances provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyBalanceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCompanyBalanceUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	companyBalanceUpsertCacheMut.RLock()
	cache, cached := companyBalanceUpsertCache[key]
	companyBalanceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			companyBalanceAllColumns,
			companyBalanceColumnsWithDefault,
			companyBalanceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			companyBalanceAllColumns,
			companyBalancePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert company_balances, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`company_balances`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `company_balances` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(companyBalanceType, companyBalanceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(companyBalanceType, companyBalanceMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for company_balances")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == companyBalanceMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(companyBalanceType, companyBalanceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for company_balances")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for company_balances")
	}

CacheNoHooks:
	if !cached {
		companyBalanceUpsertCacheMut.Lock()
		companyBalanceUpsertCache[key] = cache
		companyBalanceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CompanyBalance record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompanyBalance) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompanyBalance provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), companyBalancePrimaryKeyMapping)
	sql := "DELETE FROM `company_balances` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from company_balances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for company_balances")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q companyBalanceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no companyBalanceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from company_balances")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company_balances")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompanyBalanceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(companyBalanceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyBalancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `company_balances` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, companyBalancePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from companyBalance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company_balances")
	}

	if len(companyBalanceAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CompanyBalance) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompanyBalance(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompanyBalanceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompanyBalanceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyBalancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `company_balances`.* FROM `company_balances` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, companyBalancePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompanyBalanceSlice")
	}

	*o = slice

	return nil
}

// CompanyBalanceExists checks if the CompanyBalance row exists.
func CompanyBalanceExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `company_balances` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if company_balances exists")
	}

	return exists, nil
}

// Exists checks if the CompanyBalance row exists.
func (o *CompanyBalance) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CompanyBalanceExists(ctx, exec, o.ID)
}