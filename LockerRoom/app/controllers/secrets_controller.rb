class SecretsController < ApplicationController
  before_action :logged_in?, only: [:create, :destroy]
  before_action :correct_user, only: :destroy

  def create
    puts "!!!!"
    puts secret_params
    @secret = current_user.secrets.build(secret_params)
    if @secret.save
      flash[:success] = "Saved!"
      redirect_to "/welcome"
    else
      flash.now[:danger] = "Password need 8 symbols length"
      render 'static_pages/home'
    end
  end

  def destroy
    @secret.destroy
    flash[:success] = "Password deleted"
    redirect_to '/home'
  end

  private
    def secret_params
      params.require(:secret).permit(:domain, :password)
    end

    def correct_user
      @secret = current_user.secrets.find_by(id: params[:id])
      if @secret.nil?
        flash.now[:danger] = "Can't find this secret"
        redirect_to '/home'
      end
      puts "CORRECT"
    end
end
