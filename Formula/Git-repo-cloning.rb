# This file was generated by GoReleaser. DO NOT EDIT.
class GitRepoCloning < Formula
  desc "MakeClones to clone repos from a google sheet"
  homepage "https://github.com/omarsagoo/Git-repo-cloning"
  version "1.0.7"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/omarsagoo/Git-repo-cloning/releases/download/v1.0.7/Git-repo-cloning_1.0.7_macOS-64bit.tar.gz"
    sha256 "ca199dc42a3ce46ab54512158441452c94937bef88b5f5f7f72758eecbd89b83"
  elsif OS.linux?
  end

  def install
    bin.install "makeclones"
  end
end