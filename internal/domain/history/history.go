package history

//go:generate go run ../../../dev/gen/dto_gen.go history
//go:generate go run ../../../dev/gen/query_gen.go history

import (
	"STUOJ/internal/infrastructure/repository/dao"
	entity "STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"fmt"
	"time"

	"STUOJ/internal/domain/history/valueobject"
)

type History struct {
	Id           int64
	UserId       int64
	ProblemId    int64
	Title        valueobject.Title
	Source       valueobject.Source
	Difficulty   entity.Difficulty
	TimeLimit    float64
	MemoryLimit  int64
	Description  valueobject.Description
	Input        valueobject.Input
	Output       valueobject.Output
	SampleInput  valueobject.Input
	SampleOutput valueobject.Output
	Hint         valueobject.Description
	Operation    entity.Operation
	CreateTime   time.Time
}

func (h *History) verify() error {
	if h.UserId == 0 {
		return fmt.Errorf("用户Id不能为空")
	}
	if !entity.Operation(h.Operation).IsValid() {
		return fmt.Errorf("操作类型不合法")
	}
	if err := h.Title.Verify(); err != nil {
		return err
	}
	if err := h.Source.Verify(); err != nil {
		return err
	}
	if err := h.Description.Verify(); err != nil {
		return err
	}
	if err := h.Input.Verify(); err != nil {
		return err
	}
	if err := h.Output.Verify(); err != nil {
		return err
	}
	if err := h.SampleInput.Verify(); err != nil {
		return err
	}
	if err := h.SampleOutput.Verify(); err != nil {
		return err
	}
	if err := h.Hint.Verify(); err != nil {
		return err
	}
	return nil
}

func (h *History) toEntity() entity.History {
	return entity.History{
		Id:           uint64(h.Id),
		UserId:       uint64(h.UserId),
		ProblemId:    uint64(h.ProblemId),
		Title:        h.Title.String(),
		Source:       h.Source.String(),
		Difficulty:   h.Difficulty,
		TimeLimit:    h.TimeLimit,
		MemoryLimit:  uint64(h.MemoryLimit),
		Description:  h.Description.String(),
		Input:        h.Input.String(),
		Output:       h.Output.String(),
		SampleInput:  h.SampleInput.String(),
		SampleOutput: h.SampleOutput.String(),
		Hint:         h.Hint.String(),
		Operation:    h.Operation,
		CreateTime:   h.CreateTime,
	}
}

func (h *History) fromEntity(history entity.History) *History {
	h.Id = int64(history.Id)
	h.UserId = int64(history.UserId)
	h.ProblemId = int64(history.ProblemId)
	h.Title = valueobject.NewTitle(history.Title)
	h.Source = valueobject.NewSource(history.Source)
	h.Difficulty = history.Difficulty
	h.TimeLimit = history.TimeLimit
	h.MemoryLimit = int64(history.MemoryLimit)
	h.Description = valueobject.NewDescription(history.Description)
	h.Input = valueobject.NewInput(history.Input)
	h.Output = valueobject.NewOutput(history.Output)
	h.SampleInput = valueobject.NewInput(history.SampleInput)
	h.SampleOutput = valueobject.NewOutput(history.SampleOutput)
	h.Hint = valueobject.NewDescription(history.Hint)
	h.Operation = history.Operation
	h.CreateTime = history.CreateTime
	return h
}

func (h *History) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryId, option.OpEqual, h.Id)
	return options
}

func (h *History) Create() (int64, error) {
	h.CreateTime = time.Now()
	if err := h.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	history, err := dao.HistoryStore.Insert(h.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(history.Id), &errors.NoError
}

func (h *History) Update() error {
	var err error
	options := h.toOption()
	_, err = dao.HistoryStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := h.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.HistoryStore.Updates(h.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (h *History) Delete() error {
	options := h.toOption()
	_, err := dao.HistoryStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.HistoryStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

type Option func(*History)

func NewHistory(option ...Option) *History {
	h := &History{}
	for _, opt := range option {
		opt(h)
	}
	return h
}

func WithId(id int64) Option {
	return func(h *History) {
		h.Id = id
	}
}

func WithUserId(userId int64) Option {
	return func(h *History) {
		h.UserId = userId
	}
}

func WithProblemId(problemId int64) Option {
	return func(h *History) {
		h.ProblemId = problemId
	}
}

func WithTitle(title string) Option {
	return func(h *History) {
		h.Title = valueobject.NewTitle(title)
	}
}

func WithSource(source string) Option {
	return func(h *History) {
		h.Source = valueobject.NewSource(source)
	}
}

func WithDifficulty(difficulty entity.Difficulty) Option {
	return func(h *History) {
		h.Difficulty = difficulty
	}
}

func WithTimeLimit(timeLimit float64) Option {
	return func(h *History) {
		h.TimeLimit = timeLimit
	}
}

func WithMemoryLimit(memoryLimit int64) Option {
	return func(h *History) {
		h.MemoryLimit = memoryLimit
	}
}

func WithDescription(description string) Option {
	return func(h *History) {
		h.Description = valueobject.NewDescription(description)
	}
}

func WithInput(input string) Option {
	return func(h *History) {
		h.Input = valueobject.NewInput(input)
	}
}

func WithOutput(output string) Option {
	return func(h *History) {
		h.Output = valueobject.NewOutput(output)
	}
}

func WithSampleInput(sampleInput string) Option {
	return func(h *History) {
		h.SampleInput = valueobject.NewInput(sampleInput)
	}
}

func WithSampleOutput(sampleOutput string) Option {
	return func(h *History) {
		h.SampleOutput = valueobject.NewOutput(sampleOutput)
	}
}

func WithHint(hint string) Option {
	return func(h *History) {
		h.Hint = valueobject.NewDescription(hint)
	}
}

func WithOperation(operation entity.Operation) Option {
	return func(h *History) {
		h.Operation = operation
	}
}
