class CreateAgentImages < ActiveRecord::Migration[5.1]
  def change
    create_table :agent_images do |t|
      t.integer :agent_id
      t.string :type
      t.boolean :replaceable, default: true
      t.string :data_file_name
      t.string :data_content_type
      t.integer :data_file_size
      t.datetime :data_updated_at

      t.timestamps null: false
    end

    add_index :agent_images, :agent_id
    add_index :agent_images, [:agent_id, :type]
    add_index :agent_images, [:agent_id, :type, :replaceable]
  end
end
