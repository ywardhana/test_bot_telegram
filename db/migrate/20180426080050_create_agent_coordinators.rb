class CreateAgentCoordinators < ActiveRecord::Migration[5.1]
  def change
    create_table :agent_coordinators do |t|
      t.integer :agent_id
      t.integer :bill_amount

      t.timestamps null: false
    end

    add_index :agent_coordinators, :agent_id
  end
end
