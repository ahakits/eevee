// Code generated by eevee. DO NOT EDIT!

package model

import (
	"api/dao"
	"api/entity"
	"context"
	"encoding/json"
	"sort"
	"strconv"

	"golang.org/x/xerrors"
)

type FieldFinder interface {
	FindAll(context.Context) (*Fields, error)
	FindByDifficulties(context.Context, []int) (*Fields, error)
	FindByDifficulty(context.Context, int) (*Fields, error)
	FindByDifficultyAndLevel(context.Context, int, int) (*Fields, error)
	FindByID(context.Context, uint64) (*Field, error)
	FindByIDs(context.Context, []uint64) (*Fields, error)
	FindByLocationX(context.Context, int) (*Fields, error)
	FindByLocationXAndLocationY(context.Context, int, int) (*Field, error)
	FindByLocationXes(context.Context, []int) (*Fields, error)
	FindByName(context.Context, string) (*Field, error)
	FindByNames(context.Context, []string) (*Fields, error)
	FindByObjectNum(context.Context, int) (*Fields, error)
	FindByObjectNums(context.Context, []int) (*Fields, error)
}

type Field struct {
	*entity.Field
	fieldDAO         dao.Field
	isAlreadyCreated bool
	savedValue       entity.Field
	conv             ModelConverter
}

type Fields struct {
	values []*Field
}

type FieldsCollection []*Fields

func NewField(value *entity.Field, fieldDAO dao.Field) *Field {
	return &Field{
		Field:    value,
		fieldDAO: fieldDAO,
	}
}

func NewFields(entities entity.Fields) *Fields {
	return &Fields{values: make([]*Field, 0, len(entities))}
}

func (m *Fields) newFields(values []*Field) *Fields {
	return &Fields{values: values}
}

func (m *Fields) Each(iter func(*Field)) {
	if m == nil {
		return
	}
	for _, value := range m.values {
		iter(value)
	}
}

func (m *Fields) EachIndex(iter func(int, *Field)) {
	if m == nil {
		return
	}
	for idx, value := range m.values {
		iter(idx, value)
	}
}

func (m *Fields) EachWithError(iter func(*Field) error) error {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if err := iter(value); err != nil {
			return xerrors.Errorf("failed to iteration: %w", err)
		}
	}
	return nil
}

func (m *Fields) EachIndexWithError(iter func(int, *Field) error) error {
	if m == nil {
		return nil
	}
	for idx, value := range m.values {
		if err := iter(idx, value); err != nil {
			return xerrors.Errorf("failed to iteration: %w", err)
		}
	}
	return nil
}

func (m *Fields) Map(mapFunc func(*Field) *Field) *Fields {
	if m == nil {
		return nil
	}
	mappedValues := []*Field{}
	for _, value := range m.values {
		mappedValue := mapFunc(value)
		if mappedValue != nil {
			mappedValues = append(mappedValues, mappedValue)
		}
	}
	return m.newFields(mappedValues)
}

func (m *Fields) Any(cond func(*Field) bool) bool {
	if m == nil {
		return false
	}
	for _, value := range m.values {
		if cond(value) {
			return true
		}
	}
	return false
}

func (m *Fields) Some(cond func(*Field) bool) bool {
	return m.Any(cond)
}

func (m *Fields) IsIncluded(cond func(*Field) bool) bool {
	return m.Any(cond)
}

func (m *Fields) All(cond func(*Field) bool) bool {
	if m == nil {
		return false
	}
	for _, value := range m.values {
		if !cond(value) {
			return false
		}
	}
	return true
}

func (m *Fields) Sort(compare func(*Field, *Field) bool) {
	if m == nil {
		return
	}
	sort.Slice(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Fields) SortStable(compare func(*Field, *Field) bool) {
	if m == nil {
		return
	}
	sort.SliceStable(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Fields) Find(cond func(*Field) bool) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if cond(value) {
			return value
		}
	}
	return nil
}

func (m *Fields) Filter(filter func(*Field) bool) *Fields {
	if m == nil {
		return nil
	}
	filteredValues := []*Field{}
	for _, value := range m.values {
		if filter(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return m.newFields(filteredValues)
}

func (m *Fields) IsEmpty() bool {
	if m == nil {
		return true
	}
	if len(m.values) == 0 {
		return true
	}
	return false
}

func (m *Fields) At(idx int) *Field {
	if m == nil {
		return nil
	}
	if idx < 0 {
		return nil
	}
	if len(m.values) > idx {
		return m.values[idx]
	}
	return nil
}

func (m *Fields) First() *Field {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[0]
	}
	return nil
}

func (m *Fields) Last() *Field {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[len(m.values)-1]
	}
	return nil
}

func (m *Fields) Compact() *Fields {
	if m == nil {
		return nil
	}
	compactedValues := []*Field{}
	for _, value := range m.values {
		if value == nil {
			continue
		}
		compactedValues = append(compactedValues, value)
	}
	return m.newFields(compactedValues)
}

func (m *Fields) Add(args ...*Field) *Fields {
	if m == nil {
		return nil
	}
	for _, value := range args {
		m.values = append(m.values, value)
	}
	return m
}

func (m *Fields) Merge(args ...*Fields) *Fields {
	if m == nil {
		return nil
	}
	for _, arg := range args {
		for _, value := range arg.values {
			m.values = append(m.values, value)
		}
	}
	return m
}

func (m *Fields) Len() int {
	if m == nil {
		return 0
	}
	return len(m.values)
}

func (m *FieldsCollection) Merge() *Fields {
	if m == nil {
		return nil
	}
	if len(*m) == 0 {
		return nil
	}
	if len(*m) == 1 {
		return (*m)[0]
	}
	values := []*Field{}
	for _, collection := range *m {
		for _, value := range collection.values {
			values = append(values, value)
		}
	}
	return (*m)[0].newFields(values)
}

func (m *Field) ToJSON(ctx context.Context) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '{')
	buf = append(buf, "\"id\":"...)
	buf = strconv.AppendUint(buf, m.ID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"name\":"...)
	buf = append(buf, strconv.Quote(m.Name)...)
	buf = append(buf, ',')
	buf = append(buf, "\"locationX\":"...)
	buf = strconv.AppendInt(buf, int64(m.LocationX), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"locationY\":"...)
	buf = strconv.AppendInt(buf, int64(m.LocationY), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"objectNum\":"...)
	buf = strconv.AppendInt(buf, int64(m.ObjectNum), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"level\":"...)
	buf = strconv.AppendInt(buf, int64(m.Level), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"difficulty\":"...)
	buf = strconv.AppendInt(buf, int64(m.Difficulty), 10)
	buf = append(buf, '}')
	return buf, nil
}

func (m *Field) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	isWritten := false
	buf = append(buf, '{')
	if option.Exists("id") {
		buf = append(buf, "\"id\":"...)
		buf = strconv.AppendUint(buf, m.ID, 10)
		isWritten = true
	}
	if option.Exists("name") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"name\":"...)
		buf = append(buf, strconv.Quote(m.Name)...)
		isWritten = true
	}
	if option.Exists("location_x") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"locationX\":"...)
		buf = strconv.AppendInt(buf, int64(m.LocationX), 10)
		isWritten = true
	}
	if option.Exists("location_y") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"locationY\":"...)
		buf = strconv.AppendInt(buf, int64(m.LocationY), 10)
		isWritten = true
	}
	if option.Exists("object_num") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"objectNum\":"...)
		buf = strconv.AppendInt(buf, int64(m.ObjectNum), 10)
		isWritten = true
	}
	if option.Exists("level") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"level\":"...)
		buf = strconv.AppendInt(buf, int64(m.Level), 10)
		isWritten = true
	}
	if option.Exists("difficulty") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"difficulty\":"...)
		buf = strconv.AppendInt(buf, int64(m.Difficulty), 10)
		isWritten = true
	}
	buf = append(buf, '}')
	return buf, nil
}

func (m *Fields) ToJSON(ctx context.Context) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '[')
	for idx, value := range m.values {
		if idx != 0 {
			buf = append(buf, ',')
		}
		bytes, err := value.ToJSON(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, bytes...)
	}
	buf = append(buf, ']')
	return buf, nil
}

func (m *Fields) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '[')
	for idx, value := range m.values {
		if idx != 0 {
			buf = append(buf, ',')
		}
		bytes, err := value.ToJSONWithOption(ctx, option)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, bytes...)
	}
	buf = append(buf, ']')
	return buf, nil
}

func (m *Field) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Field) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Fields) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Fields) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Field) UnmarshalJSON(bytes []byte) error {
	var value struct {
		*entity.Field
	}
	if err := json.Unmarshal(bytes, &value); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.Field = value.Field
	return nil
}

func (m *Fields) UnmarshalJSON(bytes []byte) error {
	var values []*Field
	if err := json.Unmarshal(bytes, &values); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.values = values
	return nil
}

func (m *Field) ToMap(ctx context.Context) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := map[string]interface{}{}
	value["id"] = m.ID
	value["name"] = m.Name
	value["locationX"] = m.LocationX
	value["locationY"] = m.LocationY
	value["objectNum"] = m.ObjectNum
	value["level"] = m.Level
	value["difficulty"] = m.Difficulty
	return value, nil
}

func (m *Field) ToMapWithOption(ctx context.Context, option *RenderOption) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := map[string]interface{}{}
	if option.Exists("id") {
		value["id"] = m.ID
	}
	if option.Exists("name") {
		value["name"] = m.Name
	}
	if option.Exists("location_x") {
		value["locationX"] = m.LocationX
	}
	if option.Exists("location_y") {
		value["locationY"] = m.LocationY
	}
	if option.Exists("object_num") {
		value["objectNum"] = m.ObjectNum
	}
	if option.Exists("level") {
		value["level"] = m.Level
	}
	if option.Exists("difficulty") {
		value["difficulty"] = m.Difficulty
	}
	return value, nil
}

func (m *Fields) ToMap(ctx context.Context) ([]map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := []map[string]interface{}{}
	for _, v := range m.values {
		mapValue, err := v.ToMap(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value = append(value, mapValue)
	}
	return value, nil
}

func (m *Fields) ToMapWithOption(ctx context.Context, option *RenderOption) ([]map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := []map[string]interface{}{}
	for _, v := range m.values {
		mapValue, err := v.ToMapWithOption(ctx, option)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value = append(value, mapValue)
	}
	return value, nil
}

func (m *Field) SetConverter(conv ModelConverter) {
	m.conv = conv
}

func (m *Field) Create(ctx context.Context) error {
	if m.fieldDAO == nil {
		// for testing
		return nil
	}
	if m.isAlreadyCreated {
		return xerrors.New("this instance has already created")
	}
	if err := m.fieldDAO.Create(ctx, m.Field); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	m.savedValue = *m.Field
	m.isAlreadyCreated = true
	return nil
}

func (m *Field) Update(ctx context.Context) error {
	if m.fieldDAO == nil {
		// for testing
		return nil
	}
	isRequiredUpdate := false
	if m.savedValue.ID != m.ID {
		isRequiredUpdate = true
	}
	if m.savedValue.Name != m.Name {
		isRequiredUpdate = true
	}
	if m.savedValue.LocationX != m.LocationX {
		isRequiredUpdate = true
	}
	if m.savedValue.LocationY != m.LocationY {
		isRequiredUpdate = true
	}
	if m.savedValue.ObjectNum != m.ObjectNum {
		isRequiredUpdate = true
	}
	if m.savedValue.Level != m.Level {
		isRequiredUpdate = true
	}
	if m.savedValue.Difficulty != m.Difficulty {
		isRequiredUpdate = true
	}
	if !isRequiredUpdate {
		return nil
	}
	if err := m.fieldDAO.Update(ctx, m.Field); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	m.savedValue = *m.Field
	return nil
}

func (m *Field) Delete(ctx context.Context) error {
	if m.fieldDAO == nil {
		// for testing
		return nil
	}
	if err := m.fieldDAO.DeleteByID(ctx, m.ID); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

func (m *Field) SetAlreadyCreated(isAlreadyCreated bool) {
	m.isAlreadyCreated = isAlreadyCreated
}

func (m *Field) SetSavedValue(savedValue *entity.Field) {
	m.savedValue = *savedValue
}

func (m *Field) Save(ctx context.Context) error {
	if m.isAlreadyCreated {
		if err := m.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}
	if err := m.Create(ctx); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	return nil
}

func (m *Fields) Create(ctx context.Context) error {
	if err := m.EachWithError(func(v *Field) error {
		if err := v.Create(ctx); err != nil {
			return xerrors.Errorf("failed to Create: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Fields: %w", err)
	}
	return nil
}

func (m *Fields) Update(ctx context.Context) error {
	if err := m.EachWithError(func(v *Field) error {
		if err := v.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Fields: %w", err)
	}
	return nil
}

func (m *Fields) Save(ctx context.Context) error {
	if err := m.EachWithError(func(v *Field) error {
		if err := v.Save(ctx); err != nil {
			return xerrors.Errorf("failed to Save: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Fields: %w", err)
	}
	return nil
}

func (m *Fields) UniqueID() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.ID]; exists {
			return false
		}
		filterMap[value.ID] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByID() map[uint64]*Fields {
	if m == nil {
		return nil
	}
	values := map[uint64]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.ID]; !exists {
			values[value.ID] = &Fields{}
		}
		values[value.ID].Add(value)
	}
	return values
}

func (m *Fields) IDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.ID)
	}
	return values
}

func (m *Fields) UniqueName() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[string]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.Name]; exists {
			return false
		}
		filterMap[value.Name] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByName() map[string]*Fields {
	if m == nil {
		return nil
	}
	values := map[string]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.Name]; !exists {
			values[value.Name] = &Fields{}
		}
		values[value.Name].Add(value)
	}
	return values
}

func (m *Fields) Names() []string {
	if m == nil {
		return nil
	}
	values := []string{}
	for _, value := range m.values {
		values = append(values, value.Name)
	}
	return values
}

func (m *Fields) UniqueLocationX() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.LocationX]; exists {
			return false
		}
		filterMap[value.LocationX] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByLocationX() map[int]*Fields {
	if m == nil {
		return nil
	}
	values := map[int]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.LocationX]; !exists {
			values[value.LocationX] = &Fields{}
		}
		values[value.LocationX].Add(value)
	}
	return values
}

func (m *Fields) LocationXes() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.LocationX)
	}
	return values
}

func (m *Fields) UniqueLocationY() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.LocationY]; exists {
			return false
		}
		filterMap[value.LocationY] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByLocationY() map[int]*Fields {
	if m == nil {
		return nil
	}
	values := map[int]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.LocationY]; !exists {
			values[value.LocationY] = &Fields{}
		}
		values[value.LocationY].Add(value)
	}
	return values
}

func (m *Fields) LocationIes() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.LocationY)
	}
	return values
}

func (m *Fields) UniqueObjectNum() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.ObjectNum]; exists {
			return false
		}
		filterMap[value.ObjectNum] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByObjectNum() map[int]*Fields {
	if m == nil {
		return nil
	}
	values := map[int]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.ObjectNum]; !exists {
			values[value.ObjectNum] = &Fields{}
		}
		values[value.ObjectNum].Add(value)
	}
	return values
}

func (m *Fields) ObjectNums() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.ObjectNum)
	}
	return values
}

func (m *Fields) UniqueLevel() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.Level]; exists {
			return false
		}
		filterMap[value.Level] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByLevel() map[int]*Fields {
	if m == nil {
		return nil
	}
	values := map[int]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.Level]; !exists {
			values[value.Level] = &Fields{}
		}
		values[value.Level].Add(value)
	}
	return values
}

func (m *Fields) Levels() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.Level)
	}
	return values
}

func (m *Fields) UniqueDifficulty() *Fields {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *Field) bool {
		if _, exists := filterMap[value.Difficulty]; exists {
			return false
		}
		filterMap[value.Difficulty] = struct{}{}
		return true
	})
}

func (m *Fields) GroupByDifficulty() map[int]*Fields {
	if m == nil {
		return nil
	}
	values := map[int]*Fields{}
	for _, value := range m.values {
		if _, exists := values[value.Difficulty]; !exists {
			values[value.Difficulty] = &Fields{}
		}
		values[value.Difficulty].Add(value)
	}
	return values
}

func (m *Fields) Difficulties() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.Difficulty)
	}
	return values
}

func (m *Fields) FirstByID(a0 uint64) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByID(a0 uint64) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}

func (m *Fields) FirstByName(a0 string) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.Name != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByName(a0 string) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.Name != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}

func (m *Fields) FirstByLocationXAndLocationY(a0 int, a1 int) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.LocationX != a0 {
			continue
		}
		if value.LocationY != a1 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByLocationXAndLocationY(a0 int, a1 int) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.LocationX != a0 {
			continue
		}
		if value.LocationY != a1 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}

func (m *Fields) FirstByLocationX(a0 int) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.LocationX != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByLocationX(a0 int) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.LocationX != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}

func (m *Fields) FirstByObjectNum(a0 int) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.ObjectNum != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByObjectNum(a0 int) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.ObjectNum != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}

func (m *Fields) FirstByDifficultyAndLevel(a0 int, a1 int) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.Difficulty != a0 {
			continue
		}
		if value.Level != a1 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByDifficultyAndLevel(a0 int, a1 int) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.Difficulty != a0 {
			continue
		}
		if value.Level != a1 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}

func (m *Fields) FirstByDifficulty(a0 int) *Field {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.Difficulty != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Fields) FilterByDifficulty(a0 int) *Fields {
	if m == nil {
		return nil
	}
	values := []*Field{}
	for _, value := range m.values {
		if value.Difficulty != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newFields(values)
}