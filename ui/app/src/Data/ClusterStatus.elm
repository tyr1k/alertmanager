{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/tyr1k/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.ClusterStatus exposing (ClusterStatus, Status(..), decoder, encoder)

import Data.PeerStatus as PeerStatus exposing (PeerStatus)
import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias ClusterStatus =
    { name : Maybe String
    , status : Status
    , peers : Maybe (List PeerStatus)
    }


type Status
    = Ready
    | Settling
    | Disabled


decoder : Decoder ClusterStatus
decoder =
    Decode.succeed ClusterStatus
        |> optional "name" (Decode.nullable Decode.string) Nothing
        |> required "status" statusDecoder
        |> optional "peers" (Decode.nullable (Decode.list PeerStatus.decoder)) Nothing


encoder : ClusterStatus -> Encode.Value
encoder model =
    Encode.object
        [ ( "name", Maybe.withDefault Encode.null (Maybe.map Encode.string model.name) )
        , ( "status", statusEncoder model.status )
        , ( "peers", Maybe.withDefault Encode.null (Maybe.map (Encode.list PeerStatus.encoder) model.peers) )
        ]


statusDecoder : Decoder Status
statusDecoder =
    Decode.string
        |> Decode.andThen
            (\str ->
                case str of
                    "ready" ->
                        Decode.succeed Ready

                    "settling" ->
                        Decode.succeed Settling

                    "disabled" ->
                        Decode.succeed Disabled

                    other ->
                        Decode.fail <| "Unknown type: " ++ other
            )


statusEncoder : Status -> Encode.Value
statusEncoder model =
    case model of
        Ready ->
            Encode.string "ready"

        Settling ->
            Encode.string "settling"

        Disabled ->
            Encode.string "disabled"
