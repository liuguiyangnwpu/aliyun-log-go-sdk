package sls

func (c *TokenAutoUpdateClient) CreateConsumerGroup(project, logstore string, cg ConsumerGroup) (err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		err = c.logClient.CreateConsumerGroup(project, logstore, cg)
		if !c.processError(err) {
			return
		}
	}
	return
}

func (c *TokenAutoUpdateClient) UpdateConsumerGroup(project, logstore string, cg ConsumerGroup) (err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		err = c.logClient.UpdateConsumerGroup(project, logstore, cg)
		if !c.processError(err) {
			return
		}
	}
	return
}

func (c *TokenAutoUpdateClient) DeleteConsumerGroup(project, logstore string, cgName string) (err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		err = c.logClient.DeleteConsumerGroup(project, logstore, cgName)
		if !c.processError(err) {
			return
		}
	}
	return
}

func (c *TokenAutoUpdateClient) ListConsumerGroup(project, logstore string) (cgList []*ConsumerGroup, err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		cgList, err = c.logClient.ListConsumerGroup(project, logstore)
		if !c.processError(err) {
			return
		}
	}
	return
}

func (c *TokenAutoUpdateClient) HeartBeat(project, logstore string, cgName, consumer string, heartBeatShardIDs []int) (shardIDs []int, err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		shardIDs, err = c.logClient.HeartBeat(project, logstore, cgName, consumer, heartBeatShardIDs)
		if !c.processError(err) {
			return
		}
	}
	return
}

func (c *TokenAutoUpdateClient) UpdateCheckpoint(project, logstore string, cgName string, consumer string, shardID int, checkpoint string, forceSuccess bool) (err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		err = c.logClient.UpdateCheckpoint(project, logstore, cgName, consumer, shardID, checkpoint, forceSuccess)
		if !c.processError(err) {
			return
		}
	}
	return
}

func (c *TokenAutoUpdateClient) GetCheckpoint(project, logstore string, cgName string) (checkPointList []*ConsumerGroupCheckPoint, err error) {
	for i := 0; i < c.maxTryTimes; i ++ {
		checkPointList, err = c.logClient.GetCheckpoint(project, logstore, cgName)
		if !c.processError(err) {
			return
		}
	}
	return
}
