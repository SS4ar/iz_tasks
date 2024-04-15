Rails.application.routes.draw do
  root 'sessions#welcome'
  get 'home', to: 'static_pages#home'
#  get 'static_pages/help'
  get 'about', to: 'static_pages#about'
  resources :users, only: [:new, :create]
  resources :secrets, only: [:create, :destroy]
  get 'login', to: 'sessions#new'
  post 'login', to: 'sessions#create'
  get 'welcome', to: 'static_pages#home'#'sessions#welcome'
  get 'authorized', to: 'sessions#page_requires_login'
  get 'signup', to: 'users#new'
  get 'logout', to: 'sessions#destroy'
  delete 'logout', to: 'sessions#destroy'
#  get 'sessions/new'
#  get 'sessions/create'
#  get 'sessions/login'
#  get 'sessions/welcome'
#  get 'users/new'
#  get 'users/create'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
