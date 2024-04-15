class CreateSecrets < ActiveRecord::Migration[5.2]
  def change
    create_table :secrets do |t|
      t.text :domain
      t.text :password
      t.references :user, foreign_key: true

      t.timestamps null: false
    end
    add_index :secrets, [:user_id, :created_at]
  end
end
