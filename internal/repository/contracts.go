package repository

import (
	"context"
	"time"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindUserByID(id string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUserById(id string) error
}

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) UserRepository {
	return &PostgresUserRepository{db: db}
}

type ChatRepository interface {
	SaveMessage(msg domain.Chat) error
	ListMessages(userA, userB string) ([]domain.Chat, error)
}

type PostgresChatRepository struct {
	db *pgxpool.Pool
}

func NewPostgresChatRepository(db *pgxpool.Pool) ChatRepository {
	return &PostgresChatRepository{db: db}
}

type ContactRepository interface {
	Create(ctx context.Context, contact *domain.Contact) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Contact, error)
	Update(ctx context.Context, contact *domain.Contact) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Contact, int, error)
	Search(ctx context.Context, term string, page, pageSize int) ([]*domain.Contact, int, error)
	GetByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*domain.Contact, error)
	GetByEmail(ctx context.Context, email string) (*domain.Contact, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresContactRepository struct {
	db *pgxpool.Pool
}

func NewPostgresContactRepository(db *pgxpool.Pool) ContactRepository {
	return &PostgresContactRepository{db: db}
}

type CompanyRepository interface {
	Create(ctx context.Context, company *domain.Company) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Company, error)
	Update(ctx context.Context, company *domain.Company) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*domain.Company, error)
	SearchByName(ctx context.Context, query string, limit, offset int) ([]*domain.Company, error)
	FilterByStatus(ctx context.Context, status string, limit, offset int) ([]*domain.Company, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresCompanyRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCompanyRepository(db *pgxpool.Pool) CompanyRepository {
	return &PostgresCompanyRepository{db: db}
}

type DealRepository interface {
	Create(ctx context.Context, deal *domain.Deal) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Deal, error)
	Update(ctx context.Context, deal *domain.Deal) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Deal, int, error)
	GetByStageID(ctx context.Context, stageID uuid.UUID, limit, offset int) ([]*domain.Deal, error)
	UpdateStage(ctx context.Context, dealID, stageID uuid.UUID) error
	GetActiveDeals(ctx context.Context, limit, offset int) ([]*domain.Deal, error)
	GetWonDeals(ctx context.Context, limit, offset int) ([]*domain.Deal, error)
	GetLostDeals(ctx context.Context, limit, offset int) ([]*domain.Deal, error)
	GetByContactID(ctx context.Context, contactID uuid.UUID, limit, offset int) ([]*domain.Deal, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresDealRepository struct {
	db *pgxpool.Pool
}

func NewPostgresDealRepository(db *pgxpool.Pool) DealRepository {
	return &PostgresDealRepository{db: db}
}

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Task, int, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Task, error)
	GetByStatus(ctx context.Context, status string, limit, offset int) ([]*domain.Task, error)
	GetOverdueTasks(ctx context.Context, limit, offset int) ([]*domain.Task, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]*domain.Task, error)
	GetByRelatedEntity(ctx context.Context, relatedType string, relatedID uuid.UUID, limit, offset int) ([]*domain.Task, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresTaskRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTaskRepository(db *pgxpool.Pool) TaskRepository {
	return &PostgresTaskRepository{db: db}
}

type ActivityRepository interface {
	Create(ctx context.Context, activity *domain.Activity) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Activity, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.Activity, int, error)
	GetByContactID(ctx context.Context, contactID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error)
	GetByDealID(ctx context.Context, dealID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error)
	GetByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Activity, int, error)
	GetByType(ctx context.Context, activityType string, limit, offset int) ([]*domain.Activity, int, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresActivityRepository struct {
	db *pgxpool.Pool
}

func NewPostgresActivityRepository(db *pgxpool.Pool) ActivityRepository {
	return &PostgresActivityRepository{db: db}
}

type CallLogRepository interface {
	Create(ctx context.Context, callLog *domain.CallLog) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.CallLog, error)
	Update(ctx context.Context, callLog *domain.CallLog) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*domain.CallLog, error)
	GetByContactID(ctx context.Context, contactID uuid.UUID, limit, offset int) ([]*domain.CallLog, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.CallLog, error)
	GetByStatus(ctx context.Context, status string, limit, offset int) ([]*domain.CallLog, error)
	GetByTwilioCallSID(ctx context.Context, twilioCallSID string) (*domain.CallLog, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresCallLogRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCallLogRepository(db *pgxpool.Pool) CallLogRepository {
	return &PostgresCallLogRepository{db: db}
}

type NotificationRepository interface {
	Create(ctx context.Context, notification *domain.Notification) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Notification, error)
	Update(ctx context.Context, notification *domain.Notification) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Notification, error)
	GetUnreadByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Notification, error)
	MarkAsRead(ctx context.Context, id uuid.UUID) error
	MarkAllAsRead(ctx context.Context, userID uuid.UUID) error
	CountUnreadByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
	DeleteOldNotifications(ctx context.Context, userID uuid.UUID, olderThanDays int) error
	Count(ctx context.Context) (int64, error)
}

type PostgresNotificationRepository struct {
	db *pgxpool.Pool
}

func NewPostgresNotificationRepository(db *pgxpool.Pool) NotificationRepository {
	return &PostgresNotificationRepository{db: db}
}

type TagRepository interface {
	Create(ctx context.Context, tag *domain.Tag) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error)
	Update(ctx context.Context, tag *domain.Tag) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*domain.Tag, error)
	GetByEntityType(ctx context.Context, entityType string, limit, offset int) ([]*domain.Tag, error)
	GetByName(ctx context.Context, name, entityType string) (*domain.Tag, error)
	SearchByName(ctx context.Context, query string, limit, offset int) ([]*domain.Tag, error)
	IncrementUsage(ctx context.Context, id uuid.UUID) error
	DecrementUsage(ctx context.Context, id uuid.UUID) error
	GetMostUsed(ctx context.Context, limit int) ([]*domain.Tag, error)
	Count(ctx context.Context) (int64, error)
}

type PostgresTagRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTagRepository(db *pgxpool.Pool) TagRepository {
	return &PostgresTagRepository{db: db}
}

type PipelineStageRepository interface {
	Create(ctx context.Context, stage *domain.PipelineStage) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.PipelineStage, error)
	Update(ctx context.Context, stage *domain.PipelineStage) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, filters map[string]interface{}, page, pageSize int) ([]*domain.PipelineStage, int, error)
}

type PostgresPipelineStageRepository struct {
	db *pgxpool.Pool
}

func NewPostgresPipelineStageRepository(db *pgxpool.Pool) PipelineStageRepository {
	return &PostgresPipelineStageRepository{db: db}
}
