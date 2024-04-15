module ApplicationHelper
  def flashed(opts = {})
    flash.each do |message_type,message|
      concat(content_tag(:div, message, class: "alert alert-#{message_type}",role: "alert"))
    end
    nil
  end
end
