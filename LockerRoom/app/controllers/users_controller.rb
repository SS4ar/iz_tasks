class UsersController < ApplicationController

  skip_before_action :authorized, only: [:new, :create]

  def new
    @user = User.new
  end

  def create
    @user = User.find_by(username: params[:user][:username].downcase)
    if @user
      flash[:danger] = "Username is used"
      redirect_to '/signup'
    else
      @user = User.create(params.require(:user).permit(:username,:password))
      @secrets = @user.secrets
      session[:user_id] = @user.id
      redirect_to '/welcome'
    end
  end
end
