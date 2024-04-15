class StaticPagesController < ApplicationController
  def home
    @secret = current_user.secrets.build if logged_in?
  end

  def help
  end

  def about
    render file: "#{Rails.root}/about.html"
  end

  def robots
  end

end
