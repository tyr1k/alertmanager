{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/tyr1k/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.AlertmanagerStatus exposing (AlertmanagerStatus, decoder, encoder)

import Data.AlertmanagerConfig as AlertmanagerConfig exposing (AlertmanagerConfig)
import Data.ClusterStatus as ClusterStatus exposing (ClusterStatus)
import Data.VersionInfo as VersionInfo exposing (VersionInfo)
import DateTime exposing (DateTime)
import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias AlertmanagerStatus =
    { cluster : ClusterStatus
    , versionInfo : VersionInfo
    , config : AlertmanagerConfig
    , uptime : DateTime
    }


decoder : Decoder AlertmanagerStatus
decoder =
    Decode.succeed AlertmanagerStatus
        |> required "cluster" ClusterStatus.decoder
        |> required "versionInfo" VersionInfo.decoder
        |> required "config" AlertmanagerConfig.decoder
        |> required "uptime" DateTime.decoder


encoder : AlertmanagerStatus -> Encode.Value
encoder model =
    Encode.object
        [ ( "cluster", ClusterStatus.encoder model.cluster )
        , ( "versionInfo", VersionInfo.encoder model.versionInfo )
        , ( "config", AlertmanagerConfig.encoder model.config )
        , ( "uptime", DateTime.encoder model.uptime )
        ]
