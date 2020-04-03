#! /bin/bash

echo Generating NiFi Go client
export workspace_dir=$(pwd)
export wv_tmp_dir=${workspace_dir}/pkg
cd resources/client_gen
./generate_api_client.sh

cd ../../

echo Cleaning NiFi Go client
rm -fr ${wv_tmp_dir}/nifi/.swagger-codegen \
  && rm -fr ${wv_tmp_dir}/nifi/.gitignore \
  && rm -fr ${wv_tmp_dir}/nifi/.swagger-codegen-ignore \
  && rm -fr ${wv_tmp_dir}/nifi/.travis.yml \
  && rm -fr ${wv_tmp_dir}/swagger-codegen-cli-*.jar \
  && rm -fr ${wv_tmp_dir}/nifi.conf.json

echo Moving documentations files to root dirs
rm -fr ${wv_tmp_dir}/docs && mv ${wv_tmp_dir}/nifi/docs .
rm -fr ${wv_tmp_dir}/README.md && mv ${wv_tmp_dir}/nifi/README.md .