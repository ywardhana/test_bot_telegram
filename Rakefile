#!/usr/bin/env rake

ENV['RACK_ENV'] ||= 'development'

require 'erb'
require 'open3'
require 'yaml'
require 'dotenv'
require 'standalone_migrations'
require 'lhm'

Dotenv.load

# rake task for database migrations
StandaloneMigrations::Tasks.load_tasks

# rake task for deployment
