# frozen_string_literal: true

require_relative "lib/himekuri_go/version"

Gem::Specification.new do |spec|
  spec.name          = "himekuri_go"
  spec.version       = HimekuriGo::VERSION
  spec.authors       = ["Takayuki Kamiyama"]
  spec.email         = ["karuma.reason@gmail.com"]
  spec.summary       = %q{日めくりGoです、呼び出すと現在時刻表示します。}
  spec.description   = %q{Goでビルドしたシェルを呼びまします、現在時刻表示します。}
  spec.license       = "MIT"
  spec.executables   = %w(himekuri_go)
  spec.homepage      = "http://rubyist.hatenablog.com"
  spec.files         = `git ls-files -z`.split("\x0")
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.test_files    = spec.files.grep(%r{^(test|spec|features)/})
  spec.require_paths = ["lib"]
end
