package entity

type Agent struct {
  ID                                int64           `db:"id"`
  Deleted                           bool            `db:"deleted"`
  UserID                            int             `db:"user_id"`
  JoinedAt                          *time.Time      `db:"joined_at"`
  Ktp                               string          `db:"ktp"`
  Status                            int             `db:"status"`
  ReferrerID                        int64           `db:"referrer_id"`
  CashbackID                        int             `db:"cashback_id"`
  Reason                            string          `db:"reason"`
  ReferrerCommissionPaidAt          *time.Time      `db:"referrer_commission_paid_at"`
  KycVerifiedBy                     int             `db:"kyc_verified_by"`
  KycVerifiedAt                     *time.Time      `db:"kyc_verified_at"`
  KycRejectedBy                     int             `db:"kyc_rejected_by"`
  KycRejectedAt                     *time.Time      `db:"kyc_rejected_at"`
  KycRejectionReason                string          `db:"kyc_rejection_reason"`
  AgentType                         string          `db:"agent_type"`
  RegisterFrom                      string          `db:"register_from"`
  KycFrom                           string          `db:"kyc_from"`
  CreatedAt                         time.Time       `db:"created_at"`
  UpdatedAt                         time.Time       `db:"updated_at"`
}
