class CreateAgents < ActiveRecord::Migration[5.1]
  def change
    create_table :agents do |t|
      t.boolean :deleted, default: false
      t.integer :user_id
      t.datetime :joined_at
      t.string :ktp
      t.integer :status, default: 0
      t.integer :referrer_id
      t.integer :cashback_id
      t.string :reason
      t.datetime :referrer_commission_paid_at
      t.integer :kyc_verified_by
      t.datetime :kyc_verified_at
      t.integer :kyc_rejected_by
      t.datetime :kyc_rejected_at
      t.string :kyc_rejection_reason
      t.string :agent_type
      t.string :register_from
      t.string :kyc_from

      t.timestamps null: false
    end

    add_index :agents, :user_id
    add_index :agents, [:deleted, :joined_at]
    add_index :agents, :created_at
    add_index :agents, :updated_at
    add_index :agents, :ktp
    add_index :agents, :referrer_id
    add_index :agents, [:kyc_verified_by, :kyc_verified_at]
    add_index :agents, [:kyc_rejected_by, :kyc_rejected_at]
    add_index :agents, :register_from
    add_index :agents, :kyc_from
  end
end
