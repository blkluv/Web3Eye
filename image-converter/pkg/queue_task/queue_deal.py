#!/usr/bin/env python

import json
import os
import uuid
import urllib3
import logging
from confluent_kafka import Consumer, Producer
from pkg.model.resnet50 import Resnet50
from pkg.utils import imggetter
from pkg.utils import imgcheck
from pkg.utils import oss
from pkg.utils import config


def UUIDCheck(id) -> bool:
    try:
        uuid.UUID(id)
        return True
    except ValueError:
        return False


class VectorInfo(object):
    def __init__(self, id="", url="", vector=[], msg="", success=False) -> None:
        self.id = id
        self.url = url
        self.vector = vector
        self.msg = msg
        self.success = success


class VectorInfoEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, VectorInfo):
            return {
                'id': o.id,
                'url': o.url,
                'vector': o.vector,
                'msg': o.msg,
                'success': o.success,
            }
        return json.JSONEncoder.default(o)


def QueueDealImageURL2Vector():
    # Parse the command line.
    cConfig = {'bootstrap.servers': 'kafka-headless:9092',
               'group.id': 'python_default',
               'auto.offset.reset': 'earliest',
               'client.id': uuid.uuid4()}
    pConfig = {'bootstrap.servers': 'kafka-headless:9092'}

    # Create Consumer instance
    consumer = Consumer(cConfig)
    producer = Producer(pConfig)
    pTopic = "image-converter-output"
    cTopic = "image-converter-input"

    def reset_offset(consumer, partitions):
        consumer.assign(partitions)

    consumer.subscribe([cTopic], on_assign=reset_offset)

    # Poll for new messages from Kafka and print them.
    try:
        while True:
            vectorInfo = VectorInfo(success=False)
            vectorInfo.url = "" 
            vectorInfo.id = ""

            if vectorInfo.url == "":
                continue
    
            try:
                image_path, ok = imggetter.DownloadUrlImg(vectorInfo.url)
                if not ok:
                    vectorInfo.msg = "url format cannot parse,url: " + vectorInfo.url
                    logging.warning(vectorInfo.msg)
                    vectorInfo.success = False
                    producer.produce(pTopic, json.dumps(
                        vectorInfo, cls=VectorInfoEncoder), vectorInfo.id)
                    continue

                imgcheck.CheckImg(path=image_path)
                vectorInfo.vector = Resnet50().resnet50_extract_feat(img_path=image_path)

                file_s3_key = os.path.basename(image_path)
                # put the file to s3

                os.remove(path=image_path)

                vectorInfo.success = True
                producer.produce(pTopic, json.dumps(
                    vectorInfo, cls=VectorInfoEncoder), vectorInfo.id)
            except Exception as e:
                logging.error(e)
                continue

            # file_s3_key = os.path.basename(image_path)
            # # put the file to s3
            # ok=oss.OSS().put_object_retries(file_path=image_path,key=file_s3_key,bucket=config.minio_token_image_bucket,retries=3)
            # if ok:
            #     logging.info(f"put the file {file_s3_key} to s3")
            #     report_file_to_gen_car(id=vectorInfo.id,file_s3_key=file_s3_key)

            # os.remove(path=image_path)

            # vectorInfo.success = True
            # producer.produce(pTopic, json.dumps(
            #     vectorInfo, cls=VectorInfoEncoder), vectorInfo.id)

    except KeyboardInterrupt:
        pass
    finally:
        # Leave group and commit final offsets
        consumer.close()
        producer.poll(10000)
        producer.flush()

def report_file_to_gen_car(id:str,file_s3_key)-> bool:
    try:
        http = urllib3.PoolManager()
        data = json.dumps({'ID':id,"S3Key":file_s3_key}).encode()
        logging.info(config.gen_car_ip)
        http.request(
            method="POST",
            url=f"http://{config.gen_car_ip}:{config.gen_car_http_port}/v1/report/file",
            body=data
            )
        return True
    except:
        return False

def get_waiting_tokens()-> list:
    try:
        http = urllib3.PoolManager()
        data = json.dumps({"Conds":{"VectorState":{"Op":"eq","Value":"Waiting"}}}).encode()
        config.nft_meta_ip="127.0.0.1"
        resp=http.request(
            method="POST",
            url=f"http://{config.nft_meta_ip}:{config.nft_meta_http_port}/v1/get/tokens",
            body=data
            ).data
        return json.loads(resp)["Infos"]
    except Exception as e:
        logging.error(e)
        return []