# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20180426080054) do

  create_table "agent_coordinators", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=latin1" do |t|
    t.integer "agent_id"
    t.integer "bill_amount"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["agent_id"], name: "index_agent_coordinators_on_agent_id"
  end

  create_table "agent_images", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=latin1" do |t|
    t.integer "agent_id"
    t.string "type"
    t.boolean "replaceable", default: true
    t.string "data_file_name"
    t.string "data_content_type"
    t.integer "data_file_size"
    t.datetime "data_updated_at"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["agent_id", "type", "replaceable"], name: "index_agent_images_on_agent_id_and_type_and_replaceable"
    t.index ["agent_id", "type"], name: "index_agent_images_on_agent_id_and_type"
    t.index ["agent_id"], name: "index_agent_images_on_agent_id"
  end

  create_table "agents", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=latin1" do |t|
    t.boolean "deleted", default: false
    t.integer "user_id"
    t.datetime "joined_at"
    t.string "ktp"
    t.integer "status", default: 0
    t.integer "referrer_id"
    t.integer "cashback_id"
    t.string "reason"
    t.datetime "referrer_commission_paid_at"
    t.integer "kyc_verified_by"
    t.datetime "kyc_verified_at"
    t.integer "kyc_rejected_by"
    t.datetime "kyc_rejected_at"
    t.string "kyc_rejection_reason"
    t.string "agent_type"
    t.string "register_from"
    t.string "kyc_from"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["created_at"], name: "index_agents_on_created_at"
    t.index ["deleted", "joined_at"], name: "index_agents_on_deleted_and_joined_at"
    t.index ["ktp"], name: "index_agents_on_ktp"
    t.index ["kyc_from"], name: "index_agents_on_kyc_from"
    t.index ["kyc_rejected_by", "kyc_rejected_at"], name: "index_agents_on_kyc_rejected_by_and_kyc_rejected_at"
    t.index ["kyc_verified_by", "kyc_verified_at"], name: "index_agents_on_kyc_verified_by_and_kyc_verified_at"
    t.index ["referrer_id"], name: "index_agents_on_referrer_id"
    t.index ["register_from"], name: "index_agents_on_register_from"
    t.index ["updated_at"], name: "index_agents_on_updated_at"
    t.index ["user_id"], name: "index_agents_on_user_id"
  end

  create_table "virtual_product_agents", id: :integer, force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.boolean "deleted", default: false
    t.integer "user_id"
    t.string "ktp"
    t.integer "status", limit: 1, default: 0
    t.integer "referrer_id"
    t.integer "cashback_id"
    t.string "reason"
    t.datetime "referrer_commission_paid_at"
    t.string "kyc_verified_by"
    t.datetime "kyc_verified_at"
    t.integer "kyc_rejected_by"
    t.datetime "kyc_rejected_at"
    t.string "kyc_rejection_reason"
    t.string "agent_type", limit: 50
    t.string "register_from", limit: 20
    t.string "kyc_from", limit: 20
    t.datetime "joined_at"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

end
