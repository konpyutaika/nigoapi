#! /bin/bash

echo Generating NiFi Go client
export workspace_dir=$(pwd)
export wv_tmp_dir=${workspace_dir}/pkg
cd resources/client_gen
./generate_nifi_api_client.sh

cd ../../

echo Cleaning NiFi Go client
rm -fr ${wv_tmp_dir}/nifi/.swagger-codegen \
  && rm -fr ${wv_tmp_dir}/nifi/.gitignore \
  && rm -fr ${wv_tmp_dir}/nifi/.swagger-codegen-ignore \
  && rm -fr ${wv_tmp_dir}/nifi/.travis.yml \
  && rm -fr ${wv_tmp_dir}/swagger-codegen-cli-*.jar \
  && rm -fr ${wv_tmp_dir}/nifi.conf.json

echo Generating NiFi Registry Go client

export workspace_dir=$(pwd)
export wv_tmp_dir=${workspace_dir}/pkg
cd resources/client_gen
./generate_nifi_registry_api_client.sh

echo Cleaning NiFi Registry Go client
rm -fr ${wv_tmp_dir}/registry/.swagger-codegen \
  && rm -fr ${wv_tmp_dir}/registry/.gitignore \
  && rm -fr ${wv_tmp_dir}/registry/.swagger-codegen-ignore \
  && rm -fr ${wv_tmp_dir}/registry/.travis.yml \
  && rm -fr ${wv_tmp_dir}/swagger-codegen-cli-*.jar \
  && rm -fr ${wv_tmp_dir}/registry.conf.json