# frozen_string_literal: true

lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require 'open3'
require_relative "himekuri_go/version"

class HimekuriGoBasic
  def self.before
    # vesion info
    ruby_version = (RUBY_VERSION).to_s
    version = (HimekuriGo::VERSION).to_s
    himekuri_go = "himekuri_go-".to_s + version.to_s
    go_path= "$HOME/.goenv/shims/go".to_s
    
    golang_path = go_path + " " + "run" + " "+  "$HOME/.rbenv/versions/" + ruby_version + "/lib/ruby/gems/2.7.0/gems/" + himekuri_go + "/lib/himekuri_golang.go".to_s
    stdout_go, stderr_go, status_go = Open3.capture3(golang_path)
    
    stdout_go
  end
end