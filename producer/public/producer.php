<?php

function queue($data){
    $conf = new RdKafka\Conf();
    $conf->set('log_level', (string) LOG_WARNING);
    // $conf->set('debug', 'all');
    $conf->set('metadata.broker.list', $_ENV["ADDRESS"]);
    //$conf->set('enable.idempotence', 'true');

    $producer = new RdKafka\Producer($conf);

    $topic = $producer->newTopic(trim($_ENV["TOPIC"]));


    $topic->produce(RD_KAFKA_PARTITION_UA, 0, json_encode($data));
    $producer->poll(0);

    for ($flushRetries = 0; $flushRetries < 10; $flushRetries++) {
        $result = $producer->flush(10000);
        if (RD_KAFKA_RESP_ERR_NO_ERROR === $result) {
            break;
        }
    }
    
    if (RD_KAFKA_RESP_ERR_NO_ERROR !== $result) {
        throw new \RuntimeException('Was unable to flush, messages might be lost!');
    }
}