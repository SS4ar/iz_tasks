module SessionsHelper
  
  def log_in(user)
    session[:user_id] = user.id
  end

  def current_user
    @current_user ||= User.find_by(id: session[:user_id]) 
  end

  def forget(user)
    puts "FORGET"
    puts user
#    cookies.delete(:user_id)
#   cookies.delete(:remember_token)
  end

  def log_out
    forget(current_user)
    session.delete(:user_id)
#    @current_user = nil
  end
end
