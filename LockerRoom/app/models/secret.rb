class Secret < ApplicationRecord
  belongs_to :user
  validates :user_id, presence: true
  validates :password, presence: true, length: {minimum: 8}
end
