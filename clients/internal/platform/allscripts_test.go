package platform

//
//func TestGetSourceClientAllscripts_SyncAll(t *testing.T) {
//	t.Parallel()
//	//setup
//	testLogger := logrus.WithFields(logrus.Fields{
//		"type": "test",
//	})
//	mockCtrl := gomock.NewController(t)
//	defer mockCtrl.Finish()
//	fakeDatabase := mock_models.NewMockDatabaseRepository(mockCtrl)
//	fakeDatabase.EXPECT().UpsertRawResource(gomock.Any(), gomock.Any(), gomock.Any()).Times(158).Return(true, nil)
//
//	fakeSourceCredential := mock_models.NewMockSourceCredential(mockCtrl)
//	fakeSourceCredential.EXPECT().GetPatientId().AnyTimes().Return("6709dc13-ca3e-4969-886a-fe0889eb8256")
//	fakeSourceCredential.EXPECT().GetSourceType().AnyTimes().Return(pkg.SourceTypeAllscripts)
//	fakeSourceCredential.EXPECT().GetApiEndpointBaseUrl().AnyTimes().Return("https://pro171fmh.open.allscripts.com/open")
//
//	httpClient := base.OAuthVcrSetup(t, false)
//	client, _, err := GetSourceClientAllscripts(pkg.FastenLighthouseEnvSandbox, context.Background(), testLogger, fakeSourceCredential, httpClient)
//
//	//test
//	resp, err := client.SyncAll(fakeDatabase)
//	require.NoError(t, err)
//
//	//assert
//	require.NoError(t, err)
//	require.Equal(t, 158, resp.TotalResources)
//	require.Equal(t, 158, len(resp.UpdatedResources))
//}
