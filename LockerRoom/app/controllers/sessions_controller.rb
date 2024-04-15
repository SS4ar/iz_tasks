class SessionsController < ApplicationController

  skip_before_action :authorized, only: [:new, :create, :welcome]

  def new
  end

  def create
    @user = User.find_by(username: params[:session][:username].downcase)
    if @user && @user.authenticate(params[:session][:password])
      helpers.log_in(@user)
      redirect_to '/welcome'
    else
      flash.now[:danger] = "Invalid username or password"
      render 'new'
#      redirect_to '/login'
    end
  end

  def login
  end

  def welcome
   redirect_to '/home' unless !logged_in?
  end

  def page_requires_login
  end

  def destroy
    helpers.log_out
    @user = nil
    reset_session
    redirect_to root_url 
  end
end
